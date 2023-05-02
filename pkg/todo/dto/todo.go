package dto

import (
	"fmt"
	"github.com/pkg/errors"
	"time"
)

type CreateTodoRequest struct {
	Title           string `json:"title"`
	ActivityGroupID int    `json:"activity_group_id"`
	IsActive        bool   `json:"is_active"`
}

func (r *CreateTodoRequest) Validate() error {
	var emptyFields []string
	if r.Title == "" {
		emptyFields = append(emptyFields, "title")
	}
	if r.ActivityGroupID == 0 {
		emptyFields = append(emptyFields, "activity_group_id")
	}

	if len(emptyFields) > 0 {
		return errors.New(fmt.Sprintf("field is required %s", emptyFields))
	}

	return nil
}

type ReadTodoRequest struct {
	ActivityGroupID int
}

type ReadTodoResponse struct {
	Id              int64      `json:"id"`
	ActivityGroupId int        `json:"activity_group_id"`
	Title           string     `json:"title"`
	IsActive        bool       `json:"is_active"`
	Priority        string     `json:"priority"`
	CreatedAt       time.Time  `json:"createdAt"`
	UpdatedAt       *time.Time `json:"updatedAt"`
}
