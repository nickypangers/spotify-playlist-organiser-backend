package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

const prefix = "/api"

func InitRouter() {
	r := mux.NewRouter()

	r.HandleFunc(prefix+"/getSpotifyUser", getSpotifyUserHandler)

	http.Handle("/", r)

	http.ListenAndServe(":3030", nil)
}
