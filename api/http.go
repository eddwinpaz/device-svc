package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/eddwinpaz/device-svc/domain"
	"github.com/eddwinpaz/device-svc/service"
)

type handler struct {
	logService service.LogServiceInterface
}

func NewHandler(logService service.LogServiceInterface) WebHandler {
	return &handler{logService: logService}
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	logs, err := h.logService.GetDeviceByID(id)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(domain.LogResponse(logs, id))

}

func (h *handler) Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	l := domain.LogRequest{}

	err := json.NewDecoder(r.Body).Decode(&l)

	if err != nil {
		http.Error(w, http.StatusText(
			http.StatusBadRequest),
			http.StatusBadRequest,
		)
		return
	}

	err = h.logService.Ingest(l.Parse())

	if err != nil {
		http.Error(w, http.StatusText(
			http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}

}
