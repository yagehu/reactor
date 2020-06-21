package reactor

import (
	"encoding/json"
	"net/http"

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

	w.WriteHeader(http.StatusOK)
}
