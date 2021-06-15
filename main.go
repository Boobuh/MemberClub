package main

import (
	"log"
	"net/http"

	"github.com/Boobuh/MemberClub/web"
)

func main() {
	router := web.NewRouter()
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
	}
	log.Fatal(srv.ListenAndServe())
}
