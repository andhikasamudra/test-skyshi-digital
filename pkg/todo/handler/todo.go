package handler

import (
	"github.com/andhikasamudra/test-skyshi-digital/internal"
	"github.com/andhikasamudra/test-skyshi-digital/pkg/todo/dto"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (h *Handler) CreateTodo() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request dto.CreateTodoRequest

		err := c.BodyParser(&request)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(internal.BuildErrorResponse(err))
		}

		result, err := h.TodoService.CreateTodo(c, request)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(internal.BuildErrorResponse(err))
		}

		return c.JSON(internal.BuildResponse(result))
	}
}
func (h *Handler) ReadTodos() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request dto.ReadTodoRequest

		err := c.QueryParser(&request)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(internal.BuildErrorResponse(err))
		}

		result, err := h.TodoService.ReadTodos(c, request)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(internal.BuildErrorResponse(err))
		}

		return c.JSON(internal.BuildResponse(result))
	}
}
func (h *Handler) ReadTodo() fiber.Handler {
	return func(c *fiber.Ctx) error {

		result, err := h.TodoService.ReadTodo(c)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(internal.BuildErrorResponse(err))
		}

		return c.JSON(internal.BuildResponse(result))
	}
}
func (h *Handler) UpdateTodo() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request dto.UpdateTodoRequest

		err := c.BodyParser(&request)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(internal.BuildErrorResponse(err))
		}

		err = h.TodoService.UpdateTodo(c, request)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(internal.BuildErrorResponse(err))
		}

		return c.JSON(internal.BuildResponse(nil))
	}
}

func (h *Handler) DeleteTodo() fiber.Handler {
	return func(c *fiber.Ctx) error {

		err := h.TodoService.DeleteTodo(c)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(internal.BuildErrorResponse(err))
		}

		return c.JSON(internal.BuildResponse(nil))
	}
}
