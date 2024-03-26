package users

import (
	"github.com/ericolvr/goapi/internal/domain"
)

type Users struct {
	params *UsersParams
}

func NewUsers(opts ...Options) (*Users, error) {
	params, err := newUsersParams(opts...)
	if err != nil {
		return nil, err
	}
	return &Users{params: params}, nil
}

func (u *Users) Create(user domain.User) error {
	// if exists, return error
	// ctx := context.Background()
	// err := u.params.GetDB().CheckIfUserExists(ctx, user)
	// if err != nil {
	// 	return err
	// }

	// return u.params.GetDB().CreateUser(ctx, user)
	return nil
}

func (u *Users) Update(user domain.User) error {
	// if exists, return error
	// ctx := context.Background()
	// err := u.params.GetDB().CheckIfUserExists(ctx, user)
	// if err != nil {
	// 	return err
	// }

	// return u.params.GetDB().CreateUser(ctx, user)
	return nil
}

func (u *Users) Delete(user domain.User) error {
	// if exists, return error
	// ctx := context.Background()
	// err := u.params.GetDB().CheckIfUserExists(ctx, user)
	// if err != nil {
	// 	return err
	// }

	// return u.params.GetDB().CreateUser(ctx, user)
	return nil
}

func (u *Users) Database() string {
	// return u.params.Database()
	return ""
}

func (u *Users) SetDatabase(database string) {
	// u.params.SetDatabase(database)
}
