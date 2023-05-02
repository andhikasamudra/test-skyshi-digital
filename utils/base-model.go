package utils

import (
	"context"
	"github.com/uptrace/bun"
	"time"
)

type BaseModel struct {
	ID        int64 `bun:"id,pk,autoincrement"`
	CreatedAt time.Time
	UpdatedAt *time.Time
}

func (m *BaseModel) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	now := time.Now()
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = now
	case *bun.UpdateQuery:
		m.UpdatedAt = &now
	}
	return nil
}
