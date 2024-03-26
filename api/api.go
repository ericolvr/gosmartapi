package api

import (
	"net/http"

	"github.com/ericolvr/goapi/api/handlers"
	server "github.com/thiagozs/go-echowr"
)

type Api struct {
	params  *ApiParams
	server  *server.Server
	handler *handlers.ApiHandlers
}

func New(opts ...Options) (*Api, error) {
	params, err := newApiParams(opts...)
	if err != nil {
		return nil, err
	}

	// Server options config
	optsSrv := []server.Options{
		server.WithHost(params.Config().GetApp().RestHost),
		server.WithPort(params.Config().GetApp().RestPort),
	}

	servr, err := server.NewServer(optsSrv...)
	if err != nil {
		return nil, err
	}

	// Handlers options config
	hopts := []handlers.Options{
		handlers.WithConfig(params.Config()),
	}

	handls, err := handlers.New(hopts...)
	if err != nil {
		return nil, err
	}

	return &Api{
		params:  params,
		server:  servr,
		handler: handls,
	}, nil
}

func (a *Api) LoadRoutes() {
	// system root router
	sys := server.NewRouters()
	sys.AddRouter("/healthcheck",
		server.Methods{http.MethodGet: a.handler.Healthcheck})

	a.server.RegisterRouters(server.ROOT, sys)
}

func (a *Api) Start() {
	a.server.Start()
}

func (a *Api) GracefulShutdown() error {
	return a.server.GracefulShutdown()
}
