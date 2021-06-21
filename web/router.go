package web

import (
	"log"
	"net/http"

	"github.com/Boobuh/MemberClub/storage"
	"github.com/Boobuh/MemberClub/web/members"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func NewRouter(logger *log.Logger) *mux.Router {
	router := mux.NewRouter()
	s := storage.NewStorage()
	h := members.NewHandler(s, logger)
	router.HandleFunc("/members", h.Get).Methods(http.MethodGet)
	router.HandleFunc("/members", h.Post).Methods(http.MethodPost)
	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"content-type"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowCredentials(),
	)

	router.Use(cors)
	return router
}
