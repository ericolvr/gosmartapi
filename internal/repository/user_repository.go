package repository

import (
	"database/sql"
	"errors"

	"github.com/ericolvr/goapi/internal/domain"
)

type UserRepository interface {
	Create(user *domain.User) error
	GetUsers() ([]*domain.User, error)
	GetByID(id int64) (*domain.User, error)
	Update(user *domain.User) error
	Delete(id int64) error
	GetByDocument(document string) (*domain.User, error)
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
	query := "INSERT INTO users (name, document, Role, Password, Photo, Completed) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := r.db.Exec(query, user.Name, user.Document, user.Role, user.Password, user.Photo, user.Completed)
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

func (r *mysqlUserRepository) GetUsers() ([]*domain.User, error) {
	query := "SELECT id, name, document, role, password, photo, completed FROM users"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*domain.User
	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.ID, &user.Name, &user.Document, &user.Role, &user.Password, &user.Photo, &user.Completed)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (r *mysqlUserRepository) GetByID(id int64) (*domain.User, error) {
	query := "SELECT id, name, document, role, password, photo, completed FROM users WHERE id = ?"
	row := r.db.QueryRow(query, id)

	var user domain.User
	err := row.Scan(&user.ID, &user.Name, &user.Document, &user.Role, &user.Password, &user.Photo, &user.Completed)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
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

func (r *mysqlUserRepository) GetByDocument(document string) (*domain.User, error) {
	query := "SELECT id, name FROM users WHERE document = ?"
	row := r.db.QueryRow(query, document)

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
