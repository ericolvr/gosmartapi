package api

import (
	"github.com/ericolvr/goapi/api/handlers"
	"github.com/ericolvr/goapi/config"
	"github.com/ericolvr/goapi/internal/services/users"
)

type Options func(*ApiParams) error

type ApiParams struct {
	config   *config.Config
	handlers *handlers.ApiHandlers
	users    *users.Users
}

func newApiParams(opts ...Options) (*ApiParams, error) {
	p := &ApiParams{}
	for _, o := range opts {
		if err := o(p); err != nil {
			return nil, err
		}
	}
	return p, nil
}

func WithConfig(cfg *config.Config) Options {
	return func(p *ApiParams) error {
		p.config = cfg
		return nil
	}
}

func WithHandlers(hand *handlers.ApiHandlers) Options {
	return func(p *ApiParams) error {
		p.handlers = hand
		return nil
	}
}

func WithUsers(us *users.Users) Options {
	return func(p *ApiParams) error {
		p.users = us
		return nil
	}
}

// getters -----

func (p *ApiParams) Config() *config.Config {
	return p.config
}

func (p *ApiParams) Handlers() *handlers.ApiHandlers {
	return p.handlers
}

func (p *ApiParams) Users() *users.Users {
	return p.users
}
