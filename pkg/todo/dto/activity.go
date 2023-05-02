package dto

import (
	"fmt"

	"github.com/pkg/errors"
)

type CreateActivityRequest struct {
	Title string `json:"title"`
	Email string `json:"email"`
}

func (r *CreateActivityRequest) Validate() error {
	var emptyFields []string
	if r.Title == "" {
		emptyFields = append(emptyFields, "title")
	}
	if r.Email == "" {
		emptyFields = append(emptyFields, "email")
	}

	if len(emptyFields) > 0 {
		return errors.New(fmt.Sprintf("field is required %s", emptyFields))
	}

	return nil
}

type UpdateActivityRequest struct {
	Title string `json:"title"`
}

func (r *UpdateActivityRequest) Validate() error {
	var emptyFields []string
	if r.Title == "" {
		emptyFields = append(emptyFields, "title")
	}

	if len(emptyFields) > 0 {
		return errors.New(fmt.Sprintf("field is required %s", emptyFields))
	}

	return nil
}

type UpdateTodoRequest struct {
	Title    *string `json:"title"`
	Priority *string `json:"priority"`
	IsActive *bool   `json:"is_active"`
}
