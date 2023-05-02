package models

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type Activity struct {
	ID        int64      `bun:"id,pk,autoincrement" json:"id"`
	Title     string     `json:"title"`
	Email     string     `json:"email"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`

	bun.BaseModel `bun:"table:activities"`
}

var _ bun.BeforeAppendModelHook = (*Activity)(nil)

func (m *Activity) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	now := time.Now()
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = now
	case *bun.UpdateQuery:
		m.UpdatedAt = &now
	}
	return nil
}

type Model struct {
	db *bun.DB
}

func NewModel(db *bun.DB) *Model {
	return &Model{
		db: db,
	}
}

func (r *Model) CreateActivity(ctx context.Context, activity Activity) (*Activity, error) {
	_, err := r.db.NewInsert().Model(&activity).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return &activity, nil

}

func (r *Model) ReadActivities(ctx context.Context) ([]Activity, error) {
	var activities []Activity

	err := r.db.NewSelect().Model(&activities).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return activities, nil
}

func (r *Model) ReadActivity(ctx context.Context, id int) (*Activity, error) {
	var result Activity

	err := r.db.NewSelect().Model(&result).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *Model) UpdateActivity(ctx context.Context, activity Activity, updatedColumn []string) error {
	_, err := r.db.NewUpdate().
		Model(&activity).
		Column(updatedColumn...).
		WherePK().
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *Model) DeleteActivity(ctx context.Context, activity Activity) error {
	_, err := r.db.NewDelete().
		Model(&activity).
		WherePK().
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
