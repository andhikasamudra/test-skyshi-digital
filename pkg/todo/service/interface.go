package service

import (
	"github.com/andhikasamudra/test-skyshi-digital/pkg/todo/dto"
	"github.com/andhikasamudra/test-skyshi-digital/pkg/todo/models"
	"github.com/gofiber/fiber/v2"
)

type TodoServiceInterface interface {
	CreateActivity(ctx *fiber.Ctx, request dto.CreateActivityRequest) (*models.Activity, error)
	ReadActivities(ctx *fiber.Ctx) ([]models.Activity, error)
	ReadActivity(ctx *fiber.Ctx) (*models.Activity, error)
	UpdateActivity(ctx *fiber.Ctx, request dto.UpdateActivityRequest) error
	DeleteActivity(ctx *fiber.Ctx) error

	//Todo
	CreateTodo(ctx *fiber.Ctx, request dto.CreateTodoRequest) (*models.Todo, error)
	ReadTodos(ctx *fiber.Ctx, request dto.ReadTodoRequest) ([]dto.ReadTodoResponse, error)
	ReadTodo(ctx *fiber.Ctx) (*dto.ReadTodoResponse, error)
	UpdateTodo(ctx *fiber.Ctx, request dto.UpdateTodoRequest) error
	DeleteTodo(ctx *fiber.Ctx) error
}
