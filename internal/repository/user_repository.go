package repository

import (
	"database/sql"
	"errors"

	"github.com/ericolvr/goapi/internal/domain"
)

type UserRepository interface {
	Create(user *domain.User) error
	GetByID(id int64) (*domain.User, error)
	Update(user *domain.User) error
	Delete(id int64) error
	GetByEmail(email string) (*domain.User, error)
}

type mysqlUserRepository struct {
	db *sql.DB
}

func NewMySQLUserRepository(db *sql.DB) UserRepository {
	return &mysqlUserRepository{
		db: db,
	}
}

func (r *mysqlUserRepository) Create(user *domain.User) error {
	query := "INSERT INTO users (name, email) VALUES (?, ?)"
	result, err := r.db.Exec(query, user.Name, user.Email)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = id
	return nil
}

func (r *mysqlUserRepository) GetByID(id int64) (*domain.User, error) {
	query := "SELECT id, name FROM users WHERE id = ?"
	row := r.db.QueryRow(query, id)

	var user domain.User
	err := row.Scan(&user.ID, &user.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // User not found
		}
		return nil, err
	}

	return &user, nil
}

func (r *mysqlUserRepository) Update(user *domain.User) error {
	query := "UPDATE users SET name = ? WHERE id = ?"
	_, err := r.db.Exec(query, user.Name, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *mysqlUserRepository) Delete(id int64) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *mysqlUserRepository) GetByEmail(email string) (*domain.User, error) {
	query := "SELECT id, name FROM users WHERE email = ?"
	row := r.db.QueryRow(query, email)

	var user domain.User
	err := row.Scan(&user.ID, &user.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
