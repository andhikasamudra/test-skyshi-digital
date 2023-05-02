package todo

import (
	"github.com/andhikasamudra/test-skyshi-digital/pkg/todo/handler"
	"github.com/andhikasamudra/test-skyshi-digital/pkg/todo/models"
	"github.com/andhikasamudra/test-skyshi-digital/pkg/todo/service"
	"github.com/gofiber/fiber/v2"
	"github.com/uptrace/bun"
)

func InitRoute(r fiber.Router, db *bun.DB) {
	m := models.NewModel(db)
	s := service.NewService(service.Dependency{
		TodoModel: m,
	})
	h := handler.NewHandler(handler.Dependency{
		TodoService: s,
	})

	apiActivityGroup := r.Group("/activity-groups")
	apiActivityGroup.Post("", h.CreateActivity())
	apiActivityGroup.Get("/:id", h.ReadActivity())
	apiActivityGroup.Get("", h.ReadActivities())
	apiActivityGroup.Patch("/:id", h.UpdateActivity())
	apiActivityGroup.Delete("/:id", h.DeleteActivity())

	apiTodoGroup := r.Group("/todo-items")
	apiTodoGroup.Post("", h.CreateTodo())
	apiTodoGroup.Get("", h.ReadTodos())
	apiTodoGroup.Get("/:id", h.ReadTodo())
	apiTodoGroup.Patch("/:id", h.UpdateTodo())
	apiTodoGroup.Delete("/:id", h.DeleteTodo())
}
