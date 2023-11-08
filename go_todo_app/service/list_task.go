package service

import (
	"context"

	"github.com/shibayu36/go_todo_app/entity"
	"github.com/shibayu36/go_todo_app/store"
)

type ListTask struct {
	DB   store.Queryer
	Repo TaskLister
}

func (l *ListTask) ListTasks(ctx context.Context) (entity.Tasks, error) {
	return l.Repo.ListTasks(ctx, l.DB)
}
