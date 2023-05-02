package service

import (
	"github.com/andhikasamudra/test-skyshi-digital/pkg/todo/dto"
	"github.com/andhikasamudra/test-skyshi-digital/pkg/todo/models"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type Dependency struct {
	TodoModel models.TodoInterface
}

type TodoService struct {
	TodoModel models.TodoInterface
}

func NewService(d Dependency) *TodoService {
	return &TodoService{
		TodoModel: d.TodoModel,
	}
}

func (s *TodoService) CreateActivity(ctx *fiber.Ctx, request dto.CreateActivityRequest) (*models.Activity, error) {
	activity := models.Activity{
		Title: request.Title,
		Email: request.Email,
	}
	result, err := s.TodoModel.CreateActivity(ctx.Context(), activity)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *TodoService) ReadActivities(ctx *fiber.Ctx) ([]models.Activity, error) {
	result, err := s.TodoModel.ReadActivities(ctx.Context())
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *TodoService) ReadActivity(ctx *fiber.Ctx) (*models.Activity, error) {
	activityId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return nil, err
	}

	result, err := s.TodoModel.ReadActivity(ctx.Context(), activityId)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *TodoService) UpdateActivity(ctx *fiber.Ctx, request dto.UpdateActivityRequest) error {
	err := request.Validate()
	if err != nil {
		return err
	}

	activityId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return err
	}
	actData, err := s.TodoModel.ReadActivity(ctx.Context(), activityId)
	if err != nil {
		return err
	}

	actData.Title = request.Title
	updatedColumns := []string{"title"}

	err = s.TodoModel.UpdateActivity(ctx.Context(), *actData, updatedColumns)
	if err != nil {
		return err
	}

	return nil
}

func (s *TodoService) DeleteActivity(ctx *fiber.Ctx) error {
	activityId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return err
	}
	activity, err := s.TodoModel.ReadActivity(ctx.Context(), activityId)
	if err != nil {
		return err
	}

	err = s.TodoModel.DeleteActivity(ctx.Context(), *activity)
	if err != nil {
		return err
	}

	return nil
}
