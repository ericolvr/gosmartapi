package handlers

import (
	"github.com/ericolvr/goapi/config"
)

type Options func(*ApiHandlersParams) error

type ApiHandlersParams struct {
	config *config.Config
}

func newApiHandlersParams(opts ...Options) (*ApiHandlersParams, error) {
	p := &ApiHandlersParams{}
	for _, o := range opts {
		if err := o(p); err != nil {
			return nil, err
		}
	}
	return p, nil
}

func WithConfig(cfg *config.Config) Options {
	return func(p *ApiHandlersParams) error {
		p.config = cfg
		return nil
	}
}

// getters -----

func (p *ApiHandlersParams) Config() *config.Config {
	return p.config
}
