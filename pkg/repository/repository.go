package repository

import (
	todo "ToDoApp"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
	GetNames() ([]string, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId, listId int) (todo.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input todo.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, item todo.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todo.TodoItem, error)
	GetById(userId, itemId int) (todo.TodoItem, error)
	Delete(userId, listId int) error
	Update(userId, itemId int, input todo.UpdateItemInput) error
}

type Status interface {
	Create(description string) (int, error)
	GetAll() ([]Status, error)
	Delete(statusId int) error
	GetUsersStatuses() ([]todo.UserStatusPage, error)
	SetStatus(userId int, statusId int) error
	DropStatus(userId int) error
}

type Admin interface {
	GetRole(userId int) (todo.Role, error)
	DeleteUser(userId int) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
	Status
	Admin
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      newTodoItemPostgres(db),
		Status:        NewStatusPostgres(db),
		Admin:         NewAdminPostgres(db),
	}
}
