package service

import (
	todo "ToDoApp"
	"ToDoApp/pkg/repository"
)

type StatusService struct {
	repo repository.Status
}

func NewStatusService(repo repository.Status) *StatusService {
	return &StatusService{repo: repo}
}

func (s *StatusService) GetUsersStatuses() ([]todo.UserStatusPage, error) {
	usersStatuses, err := s.repo.GetUsersStatuses()
	if err != nil {
		return nil, err
	}

	return usersStatuses, nil
}
