package service

import (
	todo "ToDoApp"
	"ToDoApp/pkg/repository"
)

type AdminService struct {
	repo repository.Admin
}

func NewAdminService(repo repository.Admin) *AdminService {
	return &AdminService{repo: repo}
}

func (s *AdminService) GetRole(userId int) (todo.Role, error) { return s.repo.GetRole(userId) }

func (s *AdminService) DeleteUser(userId int) error { return s.repo.DeleteUser(userId) }
