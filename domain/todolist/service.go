package todolist

import (
	"context"
	dbrepo "gohx/internal/repo/db"
)

type Service struct {
	repo TaskRepo
}

type TaskRepo interface {
	GetAllTasks(ctx context.Context) ([]*dbrepo.GetAllTasksRow, error)
}

func NewService(repo TaskRepo) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllTasks(ctx context.Context) ([]*dbrepo.GetAllTasksRow, error) {
	resp, err := s.repo.GetAllTasks(ctx)
	if err != nil {
		// do some log
		return nil, err
	}

	return resp, nil
}
