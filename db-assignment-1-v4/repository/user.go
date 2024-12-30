package repository

import (
	"a21hc3NpZ25tZW50/model"
	"database/sql"
	"fmt"
)

type UserRepository interface {
	Add(user model.User) error
	CheckAvail(user model.User) error
	FetchByID(id int) (*model.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *userRepository {
	return &userRepository{db}
}
func (u *userRepository) Add(user model.User) error {
	_, err := u.db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", 
		user.Username, user.Password)
	return err
}

func (u *userRepository) CheckAvail(user model.User) error {
	row := u.db.QueryRow("SELECT COUNT(*) FROM users WHERE username = $1 AND password = $2", 
		user.Username, user.Password)
	var count int
	err := row.Scan(&count)
	if err != nil || count == 0 {
		return fmt.Errorf("user not available")
	}
	return nil
}

func (u *userRepository) FetchByID(id int) (*model.User, error) {
	row := u.db.QueryRow("SELECT id, username, password FROM users WHERE id = $1", id)

	var user model.User
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
