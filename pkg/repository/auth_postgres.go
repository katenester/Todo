package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	todo "github.com/katenester/Todo"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db}
}
func (a *AuthPostgres) CreateUser(user todo.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES($1,$2,$3) RETURNING id", usersTable)
	row := a.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (a *AuthPostgres) GetUser(username, password string) (todo.User, error) {
	var user todo.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2 ", usersTable)
	// Получаем один результат
	err := a.db.Get(&user, query, username, password)
	return user, err
}
