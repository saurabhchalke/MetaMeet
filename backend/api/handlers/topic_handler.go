package handlers

import (
	"errors"
	"fmt"
	"github.com/darmiel/perplex/api/presenter"
	"github.com/darmiel/perplex/api/services"
	"github.com/darmiel/perplex/pkg/lexorank"
	"github.com/darmiel/perplex/pkg/model"
	"github.com/darmiel/perplex/pkg/util"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	gofiberfirebaseauth "github.com/ralf-life/gofiber-firebaseauth"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type TopicHandler struct {
	srv       services.TopicService
	meetSrv   services.MeetingService
	projSrv   services.ProjectService
	userSrv   services.UserService
	logger    *zap.SugaredLogger
	validator *validator.Validate
}

func NewTopicHandler(
	srv services.TopicService,
	meetSrv services.MeetingService,
	projSrv services.ProjectService,
	userSrv services.UserService,
	logger *zap.SugaredLogger,
	validator *validator.Validate,
) *TopicHandler {
	return &TopicHandler{srv, meetSrv, projSrv, userSrv, logger, validator}
}

var ErrNoSolution = errors.New("topic requires a solution before close")

type topicDto struct {
	Title         string `validate:"required,startsnotwith= ,endsnotwith= ,min=1,max=128" json:"title"`
	Description   string `json:"description"`
	ForceSolution bool   `json:"force_solution"`
	PriorityID    uint   `json:"priority_id"`
}

func (h *TopicHandler) ValidateTopicDto(dto *topicDto) error {
	if err := h.validator.Struct(dto); err != nil {
		return err
	}
	if len(dto.Description) > MaxDescriptionLength {
		return ErrDescriptionTooLong
	}
	return nil
}

// TopicAuthorizationMiddleware is a middleware function for authorizing topic related actions.
// It fetches the relevant topic using its ID from the request params and sets it in Ctx for future use.
func (h *TopicHandler) TopicAuthorizationMiddleware(ctx *fiber.Ctx) error {
	topicID, err := ctx.ParamsInt("topic_id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(presenter.ErrorResponse(err))
	}
	m := ctx.Locals("meeting").(model.Meeting)
	topics, err := h.srv.ListTopicsForMeeting(m.ID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(presenter.ErrorResponse(err))
	}
	if topic, ok := util.Any(topics, func(t *model.Topic) bool {
		return t.ID == uint(topicID)
	}); ok {
		ctx.Locals("topic", *topic)
		return ctx.Next()
	}
	return ctx.Status(fiber.StatusUnauthorized).JSON(presenter.ErrorResponse(ErrNotFound))
}

// AddTopic adds a new topic to a meeting.
// It retrieves topic details from the request body and validates it.
func (h *TopicHandler) AddTopic(ctx *fiber.Ctx) error {
	u := ctx.Locals("user").(gofiberfirebaseauth.User)
	m := ctx.Locals("meeting").(model.Meeting)

	var payload topicDto
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(presenter.ErrorResponse(err))
	}
	if err := h.ValidateTopicDto(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(presenter.ErrorResponse(err))
	}

	topic, err := h.srv.AddTopic(u.UserID, m.ID, payload.Title, payload.Description, payload.ForceSolution, payload.PriorityID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(presenter.ErrorResponse(err))
	}

	// subscribe to topic
	if err = h.srv.SubscribeUser(topic.ID, u.UserID); err != nil {
		h.logger.Warnf("cannot subscribe user %s (creator) to topic %d: %v", u.UserID, topic.ID, err)
	}

	return ctx.Status(fiber.StatusCreated).JSON(presenter.SuccessResponse("topic created", topic))
}

// ListTopicForMeeting lists all topics for a particular meeting.
func (h *TopicHandler) ListTopicForMeeting(ctx *fiber.Ctx) error {
	m := ctx.Locals("meeting").(model.Meeting)
	topics, err := h.srv.ListTopicsForMeeting(m.ID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(presenter.ErrorResponse(err))
	}
	return ctx.Status(fiber.StatusOK).JSON(presenter.SuccessResponse("", topics))
}

func (h *TopicHandler) GetTopic(ctx *fiber.Ctx) error {
	t := ctx.Locals("topic").(model.Topic)
	if err := h.srv.Extend(&t, "Comments"); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(presenter.ErrorResponse(err))
	}
	return ctx.Status(fiber.StatusOK).JSON(presenter.SuccessResponse("topic", t))
}

// DeleteTopic deletes the topic from a meeting.
func (h *TopicHandler) DeleteTopic(ctx *fiber.Ctx) error {
	t := ctx.Locals("topic").(model.Topic)
	if err := h.srv.DeleteTopic(t.ID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(presenter.ErrorResponse(err))
	}
	return ctx.Status(fiber.StatusOK).JSON(presenter.SuccessResponse("topic deleted", nil))
}

// EditTopic edits the details of an existing topic.
func (h *TopicHandler) EditTopic(ctx *fiber.Ctx) error {
	t := ctx.Locals("topic").(model.Topic)

	var payload topicDto
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(presenter.ErrorResponse(err))
	}
	if err := h.ValidateTopicDto(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(presenter.ErrorResponse(err))
	}

	if err := h.srv.EditTopic(t.ID, payload.Title, payload.Description, payload.ForceSolution, payload.PriorityID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(presenter.ErrorResponse(err))
	}
	return ctx.Status(fiber.StatusOK).JSON(presenter.SuccessResponse("topic edited", nil))
}

// SetStatusChecked sets the status of a topic as checked (or closed).
func (h *TopicHandler) SetStatusChecked(ctx *fiber.Ctx) error {
	t := ctx.Locals("topic").(model.Topic)
	if t.ForceSolution && t.SolutionID <= 0 {
		return ctx.Status(fiber.StatusForbidden).JSON(presenter.ErrorResponse(ErrNoSolution))
	}
	if err := h.srv.CheckTopic(t.ID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(presenter.ErrorResponse(err))
	}
	return ctx.Status(fiber.StatusOK).JSON(presenter.SuccessResponse("topic closed", nil))
}

// SetStatusUnchecked sets the status of a topic as unchecked (or opened).
func (h *TopicHandler) SetStatusUnchecked(ctx *fiber.Ctx) error {
	t := ctx.Locals("topic").(model.Topic)
	if err := h.srv.UncheckTopic(t.ID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(presenter.ErrorResponse(err))
	}
	return ctx.Status(fiber.StatusOK).JSON(presenter.SuccessResponse("topic opened", nil))
}

func (h *TopicHandler) LinkTag(ctx *fiber.Ctx) error {
	topic := ctx.Locals("topic").(model.Topic)
	tag := ctx.Locals("tag").(model.Tag)
	return fiberResponseNoVal(ctx, "linked tag", h.srv.LinkTag(topic.ID, tag.ID))
}

func (h *TopicHandler) UnlinkTag(ctx *fiber.Ctx) error {
	topic := ctx.Locals("topic").(model.Topic)
	tag := ctx.Locals("tag").(model.Tag)
	return fiberResponseNoVal(ctx, "unlinked tag", h.srv.UnlinkTag(topic.ID, tag.ID))
}

func (h *TopicHandler) LinkUser(ctx *fiber.Ctx) error {
	p := ctx.Locals("project").(model.Project)
	m := ctx.Locals("meeting").(model.Meeting)
	topic := ctx.Locals("topic").(model.Topic)
	projectUser := ctx.Locals("project_user").(model.User)

	// create notification for linked user if not self link
	u := ctx.Locals("user").(gofiberfirebaseauth.User)
	if u.UserID != projectUser.ID {
		if err := h.userSrv.CreateNotification(
			projectUser.ID,
			topic.Title,
			"topic",
			"You have been assigned to a Topic",
			fmt.Sprintf("/project/%d/meeting/%d/topic/%d",
				p.ID, m.ID, topic.ID),
			"Go to Topic"); err != nil {
			h.logger.Warnf("cannot create notification for user %s: %v", projectUser.ID, err)
		}
	}

	return fiberResponseNoVal(ctx, "linked user", h.srv.LinkUser(topic.ID, projectUser.ID))
}

func (h *TopicHandler) UnlinkUser(ctx *fiber.Ctx) error {
	topic := ctx.Locals("topic").(model.Topic)
	projectUser := ctx.Locals("project_user").(model.User)
	return fiberResponseNoVal(ctx, "unlinked user", h.srv.UnlinkUser(topic.ID, projectUser.ID))
}

func (h *TopicHandler) IsSubscribed(ctx *fiber.Ctx) error {
	topic := ctx.Locals("topic").(model.Topic)
	u := ctx.Locals("user").(gofiberfirebaseauth.User)
	subscribed, err := h.srv.IsSubscribed(topic.ID, u.UserID)
	return fiberResponse(ctx, "is subscribed", subscribed, err)
}

func (h *TopicHandler) SubscribeUser(ctx *fiber.Ctx) error {
	topic := ctx.Locals("topic").(model.Topic)
	u := ctx.Locals("user").(gofiberfirebaseauth.User)
	return fiberResponseNoVal(ctx, "subscribed user", h.srv.SubscribeUser(topic.ID, u.UserID))
}

func (h *TopicHandler) UnsubscribeUser(ctx *fiber.Ctx) error {
	topic := ctx.Locals("topic").(model.Topic)
	u := ctx.Locals("user").(gofiberfirebaseauth.User)
	return fiberResponseNoVal(ctx, "unsubscribed user", h.srv.UnsubscribeUser(topic.ID, u.UserID))
}

type orderPayload struct {
	Before int `json:"before"`
	After  int `json:"after"`
}

var (
	ErrNotInSameMeeting = errors.New("topics are not in the same meeting")
)

const (
	LexoRankTop    lexorank.Rank = "aaaa"
	LexoRankBottom lexorank.Rank = "zzzz"
)

func (h *TopicHandler) UpdateOrder(ctx *fiber.Ctx) error {
	var payload orderPayload
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(presenter.ErrorResponse(err))
	}
	t := ctx.Locals("topic").(model.Topic)

	if payload.After == -1 {
		h.logger.Infof("putting topic %d on top", t.ID)

		// put topic on the top
		currentFirstTopic, err := h.srv.FindLexoRankTop(t.MeetingID)
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return ctx.Status(fiber.StatusInternalServerError).JSON(presenter.ErrorResponse(err))
			}
		} else {
			// update lexorank for top-most topic
			currentNewRank, err := LexoRankTop.Between(currentFirstTopic.LexoRank)
			if err != nil {
				return ctx.Status(fiber.StatusInternalServerError).JSON(presenter.ErrorResponse(err))
			}
			if err = h.srv.SetLexoRank(currentFirstTopic.ID, currentNewRank); err != nil {
				return ctx.Status(fiber.StatusInternalServerError).JSON(presenter.ErrorResponse(err))
			}
		}
		if err = h.srv.SetLexoRank(t.ID, LexoRankTop); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(presenter.ErrorResponse(err))
		}
		return fiberResponseNoVal(ctx, "order updated (top)", nil)
	}

	if payload.Before == -1 {
		h.logger.Infof("putting topic %d on bottom", t.ID)

		// put topic on the bottom
		currentLastTopic, err := h.srv.FindLexoRankBottom(t.MeetingID)
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return ctx.Status(fiber.StatusInternalServerError).JSON(presenter.ErrorResponse(err))
			}
		} else {
			// update lexorank for bottom-most topic
			currentNewRank, err := currentLastTopic.LexoRank.Between(LexoRankBottom)
			if err != nil {
				return ctx.Status(fiber.StatusInternalServerError).JSON(presenter.ErrorResponse(err))
			}
			if err = h.srv.SetLexoRank(currentLastTopic.ID, currentNewRank); err != nil {
				return ctx.Status(fiber.StatusInternalServerError).JSON(presenter.ErrorResponse(err))
			}
		}
		if err = h.srv.SetLexoRank(t.ID, LexoRankBottom); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(presenter.ErrorResponse(err))
		}
		return fiberResponseNoVal(ctx, "order updated (bottom)", nil)
	}

	h.logger.Infof("putting topic %d between %d and %d", t.ID, payload.Before, payload.After)

	// put topic between two topics
	// and check if in the same meeting
	topicBefore, err := h.srv.GetTopic(uint(payload.Before))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(presenter.ErrorResponse(err))
	}
	if topicBefore.MeetingID != t.MeetingID {
		return ctx.Status(fiber.StatusBadRequest).JSON(presenter.ErrorResponse(ErrNotInSameMeeting))
	}
	topicAfter, err := h.srv.GetTopic(uint(payload.After))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(presenter.ErrorResponse(err))
	}
	if topicAfter.MeetingID != t.MeetingID {
		return ctx.Status(fiber.StatusBadRequest).JSON(presenter.ErrorResponse(ErrNotInSameMeeting))
	}

	h.logger.Infof("between topics: %s <> %s", topicAfter.Title, topicBefore.Title)

	// calculate new lexorank between two topics
	newRank, err := topicAfter.LexoRank.Between(topicBefore.LexoRank)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(presenter.ErrorResponse(err))
	}
	if err = h.srv.SetLexoRank(t.ID, newRank); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(presenter.ErrorResponse(err))
	}

	h.logger.Infof("updated order for topic %d from old rank: %s to new rank: %s", t.ID, t.LexoRank, newRank)
	return ctx.Status(fiber.StatusOK).JSON(presenter.SuccessResponse("order updated (between)", nil))
}
