package models

import (
	"context"
)

type TodoInterface interface {
	CreateActivity(ctx context.Context, activity Activity) (*Activity, error)
	ReadActivities(ctx context.Context) ([]Activity, error)
	ReadActivity(ctx context.Context, id int) (*Activity, error)
	UpdateActivity(ctx context.Context, activity Activity, updatedColumn []string) error
	DeleteActivity(ctx context.Context, activity Activity) error

	// Todo
	CreateTodo(ctx context.Context, todo Todo) (*Todo, error)
	ReadTodos(ctx context.Context, filter TodoFilter) ([]Todo, error)
	ReadTodo(ctx context.Context, id int) (*Todo, error)
	UpdateTodo(ctx context.Context, todo Todo, updatedColumn []string) error
	DeleteTodo(ctx context.Context, todo Todo) error
}
