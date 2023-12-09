package routes

import (
	"github.com/darmiel/perplex/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(router fiber.Router, handler *handlers.UserHandler) {
	// list users
	router.Get("/", handler.List)
	// change username
	router.Put("/me", handler.UpdateName)
	// upcoming meetings
	router.Get("/me/upcoming-meetings", handler.UpcomingMeetings)
	router.Get("/me/search", handler.Search)
	// notifications
	router.Get("/me/notification/unread", handler.ListUnreadNotifications)
	router.Get("/me/notification/all", handler.ListAllNotifications)
	router.Delete("/me/notification/:notification_id", handler.MarkNotificationAsRead)
	router.Delete("/me/notification", handler.MarkAllNotificationsAsRead)
	// get username
	router.Get("/resolve/:user_id", handler.Resolve)
}
