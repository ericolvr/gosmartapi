package usecase

import (
	"fmt"
	"net/http"

	"github.com/ericolvr/goapi/internal/domain"
	"github.com/ericolvr/goapi/internal/repository"
)

type UserUsecase interface {
	CreateUser(user *domain.User) error
	GetUserByID(id int64) (*domain.User, error)
	UpdateUser(user *domain.User) error
	DeleteUser(id int64) error
}

type userUsecase struct {
	userRepo repository.UserRepository
}

type DuplicateEmailError struct {
	Email      string
	StatusCode int
}

func (e *DuplicateEmailError) Error() string {
	return fmt.Sprintf("user with email %s already exists", e.Email)
}

func (e *DuplicateEmailError) Status() int {
	return e.StatusCode
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (uc *userUsecase) CreateUser(user *domain.User) error {
	existingUser, err := uc.userRepo.GetByEmail(user.Email)

	if err != nil {
		return err
	}

	if existingUser != nil {
		return &DuplicateEmailError{
			Email:      user.Email,
			StatusCode: http.StatusConflict,
		}
	}

	return uc.userRepo.Create(user)
}

func (uc *userUsecase) GetUserByID(id int64) (*domain.User, error) {
	return uc.userRepo.GetByID(id)
}

func (uc *userUsecase) UpdateUser(user *domain.User) error {
	return uc.userRepo.Update(user)
}

func (uc *userUsecase) DeleteUser(id int64) error {
	return uc.userRepo.Delete(id)
}
