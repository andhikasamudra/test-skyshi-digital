package models

import (
	"context"
	"github.com/uptrace/bun"
	"time"
)

type Todo struct {
	ID              int64      `bun:"id,pk,autoincrement" json:"id"`
	Title           string     `json:"title"`
	Priority        string     `json:"priority"`
	IsActive        bool       `json:"isActive"`
	CreatedAt       time.Time  `json:"createdAt"`
	UpdatedAt       *time.Time `json:"updatedAt"`
	ActivityGroupID int        `json:"activityGroupID"`

	ActivityGroup Activity `bun:"rel:belongs-to,join:activity_group_id=id"`

	bun.BaseModel `bun:"table:todos"`
}

var _ bun.BeforeAppendModelHook = (*Todo)(nil)

func (m *Todo) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	now := time.Now()
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = now
	case *bun.UpdateQuery:
		m.UpdatedAt = &now
	}
	return nil
}

type TodoFilter struct {
	ActivityGroupID int
}

func (r *Model) CreateTodo(ctx context.Context, todo Todo) (*Todo, error) {
	_, err := r.db.NewInsert().
		Model(&todo).
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}
func (r *Model) ReadTodos(ctx context.Context, filter TodoFilter) ([]Todo, error) {
	var todos []Todo

	query := r.db.NewSelect().
		Model(&todos).
		Relation("ActivityGroup")

	if filter.ActivityGroupID != 0 {
		query.Where("activity_group_id = ?", filter.ActivityGroupID)
	}

	err := query.Scan(ctx)
	if err != nil {
		return nil, err
	}

	return todos, nil
}
func (r *Model) ReadTodo(ctx context.Context, id int) (*Todo, error) {
	var result Todo

	err := r.db.NewSelect().Model(&result).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
func (r *Model) UpdateTodo(ctx context.Context, todo Todo, updatedColumn []string) error {
	_, err := r.db.NewUpdate().
		Model(&todo).
		Column(updatedColumn...).
		WherePK().
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
func (r *Model) DeleteTodo(ctx context.Context, todo Todo) error {
	_, err := r.db.NewDelete().
		Model(&todo).
		WherePK().
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
