package reactor

import (
	"encoding/json"
	"net/http"

	reactorcontroller "github.com/yagehu/reactor/internal/controller/reactor"
	"github.com/yagehu/reactor/internal/entity"
)

type WatchEventRequest map[string][]string

func (h *handler) WatchEvent(w http.ResponseWriter, r *http.Request) {
	var req WatchEventRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	services := make([]entity.Service, 0, len(req))

	for name, tags := range req {
		tagsMap := make(map[string]struct{})

		for _, tag := range tags {
			tagsMap[tag] = struct{}{}
		}

		services = append(services, entity.NewService(
			name,
			tagsMap,
		))
	}

	_, err := h.reactorController.WatchEvent(
		r.Context(),
		&reactorcontroller.WatchEventParams{},
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
