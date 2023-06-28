package main

import (
	"time"

	"github.com/uptrace/bun"
)

type Todo struct {
	bun.BaseModel `bun:"table:todos,alias:t"`

	ID        int64     `bun:"id,pk,autoincrement" form:"id"`
	Content   string    `bun:"content,notnull" form:"content"`
	Done      bool      `bun:"done" form:"done"`
	Until     time.Time `bun:"until,nullzero" form:"until"`
	CreatedAt time.Time
	UpdatedAt time.Time `bun:",nullzero"`
	DeletedAt time.Time `bun:",soft_delete,nullzero"`
}
