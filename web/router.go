package web

import (
	"github.com/Boobuh/MemberClub/web/members"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", members.HomeHandler)
	return router
}
