package handlers

import (
	"database/sql"
	"fmt"
	"github.com/darmiel/perplex/api/presenter"
	"github.com/darmiel/perplex/api/services"
	"github.com/darmiel/perplex/pkg/model"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	gofiberfirebaseauth "github.com/ralf-life/gofiber-firebaseauth"
	"go.uber.org/zap"
	"time"
)

type ActionHandler struct {
	srv        services.ActionService
	topicSrv   services.TopicService
	meetingSrv services.MeetingService
	userSrv    services.UserService
	logger     *zap.SugaredLogger
	validator  *validator.Validate
}

func NewActionHandler(
	srv services.ActionService,
	topicSrv services.TopicService,
	meetingSrv services.MeetingService,
	userSrv services.UserService,
	logger *zap.SugaredLogger,
	validator *validator.Validate,
) *ActionHandler {
	return &ActionHandler{srv, topicSrv, meetingSrv, userSrv, logger, validator}
}

func (a ActionHandler) ListActionsForProject(ctx *fiber.Ctx) error {
	p := ctx.Locals("project").(model.Project)
	actions, err := a.srv.FindActionsByProject(p.ID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(presenter.ErrorResponse(err))
	}
	return ctx.Status(fiber.StatusOK).JSON(presenter.SuccessResponse("actions by project", actions))
}

type actionDto struct {
	Title       string `json:"title" validate:"required,min=1,max=64"`
	Description string `json:"description"`
	DueDate     string `json:"due_date" validate:"omitempty,datetime=2006-01-02T15:04:05Z07:00"`
	PriorityID  uint   `json:"priority_id"`
}

func (a ActionHandler) ValidateActionDto(dto *actionDto) (dueDate sql.NullTime, err error) {
	if err = a.validator.Struct(dto); err != nil {
		return
	}
	if len(dto.Description) > MaxDescriptionLength {
		err = ErrDescriptionTooLong
		return
	}
	if dto.DueDate != "" {
		var dueTime time.Time
		if dueTime, err = time.Parse(time.RFC3339, dto.DueDate); err != nil {
			return
		}
		dueDate = sql.NullTime{Time: dueTime, Valid: true}
	}
	return
}

func fiberResponseNoVal(ctx *fiber.Ctx, message string, err error) error {
	return fiberResponse[*bool](ctx, message, nil, err)
}

func fiberResponse[T any](ctx *fiber.Ctx, message string, value T, err error) error {
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(presenter.ErrorResponse(err))
	}
	return ctx.Status(fiber.StatusOK).JSON(presenter.SuccessResponse(message, value))
}

func (a ActionHandler) CreateAction(ctx *fiber.Ctx) error {
	p := ctx.Locals("project").(model.Project)
	u := ctx.Locals("user").(gofiberfirebaseauth.User)
	var dto actionDto
	if err := ctx.BodyParser(&dto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(presenter.ErrorResponse(err))
	}
	dueDate, err := a.ValidateActionDto(&dto)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(presenter.ErrorResponse(err))
	}
	// create action
	action, err := a.srv.CreateAction(dto.Title, dto.Description, dueDate, dto.PriorityID, p.ID, u.UserID)
	return fiberResponse(ctx, "created action", action, err)
}

func (a ActionHandler) ListActionsForProjectAndUser(ctx *fiber.Ctx) error {
	u := ctx.Locals("user").(gofiberfirebaseauth.User)
	p := ctx.Locals("project").(model.Project)
	openOnly := ctx.Query("open", "false") != "false"
	a.logger.Infof("list actions for project %d and user %s (open only: %v q/%v)", p.ID, u.UserID, openOnly, ctx.Query("open", "false"))
	actions, err := a.srv.FindActionsByProjectAndUser(p.ID, u.UserID, openOnly)
	return fiberResponse(ctx, "open actions by project and user", actions, err)
}

// :action_id

func (a ActionHandler) ActionLocalsMiddleware(ctx *fiber.Ctx) error {
	p := ctx.Locals("project").(model.Project)
	actionID, err := ctx.ParamsInt("action_id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(presenter.ErrorResponse(err))
	}
	action, err := a.srv.FindAction(uint(actionID))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(presenter.ErrorResponse(err))
	}
	if action.ProjectID != p.ID {
		return ctx.Status(fiber.StatusUnauthorized).JSON(presenter.ErrorResponse(ErrNotFound))
	}
	ctx.Locals("action", *action)
	return ctx.Next()
}

func (a ActionHandler) FindAction(ctx *fiber.Ctx) error {
	action := ctx.Locals("action").(model.Action)
	return ctx.Status(fiber.StatusOK).JSON(presenter.SuccessResponse("found action", action))
}

func (a ActionHandler) EditAction(ctx *fiber.Ctx) error {
	action := ctx.Locals("action").(model.Action)
	var dto actionDto
	if err := ctx.BodyParser(&dto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(presenter.ErrorResponse(err))
	}
	dueDate, err := a.ValidateActionDto(&dto)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(presenter.ErrorResponse(err))
	}
	// edit action
	return fiberResponseNoVal(ctx, "created action",
		a.srv.EditAction(action.ID, dto.Title, dto.Description, dueDate, dto.PriorityID))
}

func (a ActionHandler) DeleteAction(ctx *fiber.Ctx) error {
	action := ctx.Locals("action").(model.Action)
	return fiberResponseNoVal(ctx, "deleted action", a.srv.DeleteAction(action.ID))
}

func (a ActionHandler) ListActionsForMeeting(ctx *fiber.Ctx) error {
	m := ctx.Locals("meeting").(model.Meeting)

	// find topics in meeting
	topics, err := a.topicSrv.ListTopicsForMeeting(m.ID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(presenter.ErrorResponse(err))
	}

	// find all actions which are linked to one of the topics
	actionsByMeeting := make(map[uint]model.Action)
	for _, t := range topics {
		actions, err := a.srv.FindActionsByTopic(t.ID)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(presenter.ErrorResponse(err))
		}
		for _, a := range actions {
			actionsByMeeting[a.ID] = a
		}
	}

	actions := make([]model.Action, 0, len(actionsByMeeting))
	for _, a := range actionsByMeeting {
		actions = append(actions, a)
	}
	return fiberResponse(ctx, "actions by meeting", actions, err)
}

// TopicLocalsMiddleware checks if the topic exists and belongs to the project.
// It also adds the topic and meeting to the context.
func (a ActionHandler) TopicLocalsMiddleware(ctx *fiber.Ctx) error {
	p := ctx.Locals("project").(model.Project)
	topicID, err := ctx.ParamsInt("topic_id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(presenter.ErrorResponse(err))
	}
	topic, err := a.topicSrv.GetTopic(uint(topicID))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(presenter.ErrorResponse(err))
	}
	ctx.Locals("topic", *topic)
	meeting, err := a.meetingSrv.GetMeeting(topic.MeetingID)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(presenter.ErrorResponse(err))
	}
	ctx.Locals("meeting", *meeting)
	if meeting.ProjectID != p.ID {
		return ctx.Status(fiber.StatusUnauthorized).JSON(presenter.ErrorResponse(ErrNotFound))
	}
	return ctx.Next()
}

func (a ActionHandler) ListActionsForTopic(ctx *fiber.Ctx) error {
	topic := ctx.Locals("topic").(model.Topic)
	actions, err := a.srv.FindActionsByTopic(topic.ID)
	return fiberResponse(ctx, "actions by topic", actions, err)
}

func (a ActionHandler) LinkTopic(ctx *fiber.Ctx) error {
	action := ctx.Locals("action").(model.Action)
	topic := ctx.Locals("topic").(model.Topic)
	return fiberResponseNoVal(ctx, "linked topic", a.srv.LinkTopic(action.ID, topic.ID))
}

func (a ActionHandler) UnlinkTopic(ctx *fiber.Ctx) error {
	action := ctx.Locals("action").(model.Action)
	topic := ctx.Locals("topic").(model.Topic)
	return fiberResponseNoVal(ctx, "unlinked topic", a.srv.UnlinkTopic(action.ID, topic.ID))
}

// :action_id/user/:user_id

func (a ActionHandler) LinkUser(ctx *fiber.Ctx) error {
	action := ctx.Locals("action").(model.Action)
	projectUser := ctx.Locals("project_user").(model.User)
	u := ctx.Locals("user").(gofiberfirebaseauth.User)

	// create notification for linked user
	if u.UserID != projectUser.ID {
		if err := a.userSrv.CreateNotification(
			projectUser.ID,
			action.Title,
			"action",
			"you have been assigned to an action",
			fmt.Sprintf("/project/%d/action/%d", action.ProjectID, action.ID),
			"Go to Action"); err != nil {
			a.logger.Warnf("cannot create notification for user %s: %v", projectUser.ID, err)
		}
	}

	return fiberResponseNoVal(ctx, "linked user", a.srv.LinkUser(action.ID, projectUser.ID))
}

func (a ActionHandler) UnlinkUser(ctx *fiber.Ctx) error {
	action := ctx.Locals("action").(model.Action)
	projectUser := ctx.Locals("project_user").(model.User)
	return fiberResponseNoVal(ctx, "unlinked user", a.srv.UnlinkUser(action.ID, projectUser.ID))
}

// :action_id/tag/:tag_id

func (a ActionHandler) LinkTag(ctx *fiber.Ctx) error {
	action := ctx.Locals("action").(model.Action)
	tag := ctx.Locals("tag").(model.Tag)
	return fiberResponseNoVal(ctx, "linked tag", a.srv.LinkTag(action.ID, tag.ID))
}

func (a ActionHandler) UnlinkTag(ctx *fiber.Ctx) error {
	action := ctx.Locals("action").(model.Action)
	tag := ctx.Locals("tag").(model.Tag)
	return fiberResponseNoVal(ctx, "unlinked tag", a.srv.UnlinkTag(action.ID, tag.ID))
}

// :action_id/close

func (a ActionHandler) CloseAction(ctx *fiber.Ctx) error {
	action := ctx.Locals("action").(model.Action)
	return fiberResponseNoVal(ctx, "closed action", a.srv.CloseAction(action.ID))
}

func (a ActionHandler) OpenAction(ctx *fiber.Ctx) error {
	action := ctx.Locals("action").(model.Action)
	return fiberResponseNoVal(ctx, "opened action", a.srv.OpenAction(action.ID))
}
