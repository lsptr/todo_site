package service

import (
	todo "ToDoApp"
	"ToDoApp/pkg/repository"
	"fmt"
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
	for _, status := range usersStatuses {
		fmt.Println("Id: ", status.Id, "Name: ", status.Name, " Status: ", status.Status)
	}
	return usersStatuses, nil
}
