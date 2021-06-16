package members

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Boobuh/MemberClub/storage"
)

type Storage interface {
	GetAll() storage.Members
	Save(member storage.Member) error
}
type Handler struct {
	logger *log.Logger
	store  Storage
}

func NewHandler(store Storage, logger *log.Logger) *Handler {
	return &Handler{store: store, logger: logger}
}
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	h.logger.Print("new get request")
	members := h.store.GetAll()
	payload, err := json.Marshal(members)
	if err != nil {
		h.logger.Printf("error in GET members call - can't marshal object from db:%s", err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
	w.WriteHeader(http.StatusOK)
}
