package reactor

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/fx"

	"github.com/yagehu/reactor/internal/controller/reactor"
)

type Handler interface {
	WatchEvent(w http.ResponseWriter, r *http.Request)
}

type Params struct {
	fx.In

	Lifecycle         fx.Lifecycle
	Router            *mux.Router
	ReactorController reactor.Controller
}

func Register(p Params) {
	h := handler{
		reactorController: p.ReactorController,
	}

	p.Lifecycle.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			p.Router.HandleFunc("/", h.WatchEvent).Methods(http.MethodPost)

			return nil
		},
		OnStop: nil,
	})
}

type handler struct {
	reactorController reactor.Controller
}
