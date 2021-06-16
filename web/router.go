package web

import (
	"bytes"
	"log"
	"net/http"

	"github.com/Boobuh/MemberClub/storage"
	"github.com/Boobuh/MemberClub/web/members"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	s := storage.NewStorage()
	logger := log.New(new(bytes.Buffer), "testLogger: ", log.Lshortfile)
	h := members.NewHandler(s, logger)
	router.HandleFunc("/members", h.Get).Methods(http.MethodGet)
	return router
}
