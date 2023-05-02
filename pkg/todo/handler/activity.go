package handler

import (
	"net/http"

	"github.com/andhikasamudra/test-skyshi-digital/internal"
	"github.com/andhikasamudra/test-skyshi-digital/pkg/todo/dto"
	"github.com/andhikasamudra/test-skyshi-digital/pkg/todo/service"
	"github.com/gofiber/fiber/v2"
)

type Dependency struct {
	TodoService service.TodoServiceInterface
}

type Handler struct {
	TodoService service.TodoServiceInterface
}

func NewHandler(d Dependency) *Handler {
	return &Handler{
		TodoService: d.TodoService,
	}
}

func (h *Handler) CreateActivity() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request dto.CreateActivityRequest

		err := c.BodyParser(&request)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(internal.BuildErrorResponse(err))
		}

		result, err := h.TodoService.CreateActivity(c, request)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(internal.BuildErrorResponse(err))
		}

		return c.JSON(internal.BuildResponse(result))
	}
}

func (h *Handler) ReadActivities() fiber.Handler {
	return func(c *fiber.Ctx) error {

		result, err := h.TodoService.ReadActivities(c)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(internal.BuildErrorResponse(err))
		}

		return c.JSON(internal.BuildResponse(result))
	}
}
func (h *Handler) ReadActivity() fiber.Handler {
	return func(c *fiber.Ctx) error {

		result, err := h.TodoService.ReadActivity(c)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(internal.BuildErrorResponse(err))
		}

		return c.JSON(internal.BuildResponse(result))
	}
}

func (h *Handler) UpdateActivity() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request dto.UpdateActivityRequest

		err := c.BodyParser(&request)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(internal.BuildErrorResponse(err))
		}

		err = h.TodoService.UpdateActivity(c, request)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(internal.BuildErrorResponse(err))
		}

		return c.JSON(internal.BuildResponse(nil))
	}
}

func (h *Handler) DeleteActivity() fiber.Handler {
	return func(c *fiber.Ctx) error {

		err := h.TodoService.DeleteActivity(c)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(internal.BuildErrorResponse(err))
		}

		return c.JSON(internal.BuildResponse(nil))
	}
}
