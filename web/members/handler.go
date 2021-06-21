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
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	//log for members quantity if needed
	//h.logger.Printf("get call - members.length = %d", len(members))
	w.Write(payload)
	w.WriteHeader(http.StatusOK)
}
func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	h.logger.Print("new post request")
	member := storage.Member{}
	err := json.NewDecoder(r.Body).Decode(&member)
	if err != nil {
		h.logger.Printf("error in POST members call - can't decode object from request:%s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.store.Save(member)
	if err != nil {
		if err == storage.ErrInvalidEmail || err == storage.ErrInvalidName {
			h.logger.Printf("error in POST members call - invalid input:%s", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))

			return
		}
		h.logger.Printf("error in POST members call - failed to save member:%s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))

		return
	}

	w.Header().Set("Content-Type", "application/json")
	//log for members quantity if needed
	//members := h.store.GetAll()
	//h.logger.Printf("post call - members.length = %d", len(members))
	w.Write([]byte(`{"Message":"created"}`))
	w.WriteHeader(http.StatusCreated)
}
