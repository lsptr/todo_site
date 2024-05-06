package service

import (
	todo "ToDoApp"
	"ToDoApp/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
	GetIdName(username, password string) (string, int, error)
	GetAllNames() ([]string, error)
	//CheckName(name string) (bool, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId, listId int) (todo.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input todo.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, listId int, list todo.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todo.TodoItem, error)
	GetById(userId, itemId int) (todo.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input todo.UpdateItemInput) error
}

type Status interface {
	GetUsersStatuses() ([]todo.UserStatusPage, error)
}

type Admin interface {
	GetRole(userId int) (todo.Role, error)
	DeleteUser(UserId int) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
	Status
	Admin
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
		Status:        NewStatusService(repos.Status),
		Admin:         NewAdminService(repos.Admin),
	}
}
