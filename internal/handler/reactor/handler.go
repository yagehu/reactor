package reactor

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/fx"
)

type Handler interface {
	WatchEvent(w http.ResponseWriter, r *http.Request)
}

type Params struct {
	fx.In

	Lifecycle fx.Lifecycle
	Router    *mux.Router
}

func Register(p Params) {
	h := handler{}

	p.Lifecycle.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			p.Router.HandleFunc("/", h.WatchEvent).Methods(http.MethodPost)

			return nil
		},
		OnStop: nil,
	})
}

type handler struct {
}
