package users

import "github.com/ericolvr/goapi/internal/adapter/database/postgres"

type Options func(*UsersParams) error

type UsersParams struct {
	database string
	postgres *postgres.Postgres
}

func WithDatabase(database string) Options {
	return func(params *UsersParams) error {
		params.database = database
		return nil
	}
}

func WithPostgres(pg *postgres.Postgres) Options {
	return func(params *UsersParams) error {
		params.postgres = pg
		return nil
	}
}

func newUsersParams(opts ...Options) (*UsersParams, error) {
	params := &UsersParams{}
	for _, opt := range opts {
		if err := opt(params); err != nil {
			return nil, err
		}
	}
	return params, nil
}

// getters ----

func (p *UsersParams) GetDB() string {
	return p.database
}

func (p *UsersParams) GetPostgres() *postgres.Postgres {
	return p.postgres
}

// setters ----

func (p *UsersParams) SetDB(database string) {
	p.database = database
}

func (p *UsersParams) SetPostgres(pg *postgres.Postgres) {
	p.postgres = pg
}
