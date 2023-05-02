package service

import (
	"github.com/andhikasamudra/test-skyshi-digital/pkg/todo/dto"
	"github.com/andhikasamudra/test-skyshi-digital/pkg/todo/models"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func (s *TodoService) CreateTodo(ctx *fiber.Ctx, request dto.CreateTodoRequest) (*models.Todo, error) {
	err := request.Validate()
	if err != nil {
		return nil, err
	}

	_, err = s.TodoModel.ReadActivity(ctx.Context(), request.ActivityGroupID)
	if err != nil {
		return nil, err
	}

	todo := models.Todo{
		Title:           request.Title,
		IsActive:        request.IsActive,
		ActivityGroupID: request.ActivityGroupID,
	}
	result, err := s.TodoModel.CreateTodo(ctx.Context(), todo)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *TodoService) ReadTodos(ctx *fiber.Ctx, request dto.ReadTodoRequest) ([]dto.ReadTodoResponse, error) {
	result, err := s.TodoModel.ReadTodos(ctx.Context(), models.TodoFilter{
		ActivityGroupID: request.ActivityGroupID,
	})
	if err != nil {
		return nil, err
	}

	var responses []dto.ReadTodoResponse

	for _, item := range result {
		responses = append(responses, dto.ReadTodoResponse{
			Id:              item.ID,
			ActivityGroupId: item.ActivityGroupID,
			Title:           item.Title,
			IsActive:        item.IsActive,
			Priority:        item.Priority,
			CreatedAt:       item.CreatedAt,
			UpdatedAt:       item.UpdatedAt,
		})
	}

	return responses, nil
}
func (s *TodoService) ReadTodo(ctx *fiber.Ctx) (*dto.ReadTodoResponse, error) {
	todoId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return nil, err
	}

	result, err := s.TodoModel.ReadTodo(ctx.Context(), todoId)
	if err != nil {
		return nil, err
	}

	return &dto.ReadTodoResponse{
		Id:              result.ID,
		ActivityGroupId: result.ActivityGroupID,
		Title:           result.Title,
		IsActive:        result.IsActive,
		Priority:        result.Priority,
		CreatedAt:       result.CreatedAt,
		UpdatedAt:       result.UpdatedAt,
	}, nil
}
func (s *TodoService) UpdateTodo(ctx *fiber.Ctx, request dto.UpdateTodoRequest) error {
	todoID, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return err
	}

	todoData, err := s.TodoModel.ReadTodo(ctx.Context(), todoID)
	if err != nil {
		return err
	}

	var updatedColumns []string

	if request.Title != nil {
		todoData.Title = *request.Title
		updatedColumns = append(updatedColumns, "title")
	}

	if request.IsActive != nil {
		todoData.IsActive = *request.IsActive
		updatedColumns = append(updatedColumns, "is_active")
	}

	if request.Priority != nil {
		todoData.Priority = *request.Priority
		updatedColumns = append(updatedColumns, "priority")
	}

	err = s.TodoModel.UpdateTodo(ctx.Context(), *todoData, updatedColumns)
	if err != nil {
		return err
	}

	return nil
}

func (s *TodoService) DeleteTodo(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return err
	}
	todo, err := s.TodoModel.ReadTodo(ctx.Context(), id)
	if err != nil {
		return err
	}

	err = s.TodoModel.DeleteTodo(ctx.Context(), *todo)
	if err != nil {
		return err
	}

	return nil
}
