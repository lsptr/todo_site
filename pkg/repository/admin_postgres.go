package repository

import (
	todo "ToDoApp"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AdminPostgres struct {
	db *sqlx.DB
}

func NewAdminPostgres(db *sqlx.DB) *AdminPostgres {
	return &AdminPostgres{db: db}
}

func (r *AdminPostgres) GetRole(userId int) (todo.Role, error) {
	var role todo.Role
	query := fmt.Sprintf("SELECT r.id, r.title, r.description FROM %s ur INNER JOIN %s r ON ur.role_id = r.id WHERE ur.user_id = $1", usersRolesTable, rolesTable)
	err := r.db.Get(&role, query, userId)
	return role, err
}
