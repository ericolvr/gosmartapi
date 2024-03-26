package handlers

import (
	"net/http"

	server "github.com/thiagozs/go-echowr"
)

type ApiHandlers struct {
	params *ApiHandlersParams
}

func New(opts ...Options) (*ApiHandlers, error) {
	params, err := newApiHandlersParams(opts...)
	if err != nil {
		return nil, err
	}

	return &ApiHandlers{
		params: params,
	}, nil
}

// Healthcheck returns the health status of the service
func (a *ApiHandlers) Healthcheck(s server.Context) error {
	return s.JSON(http.StatusOK, map[string]interface{}{
		"status": "Healthy",
	})
}
