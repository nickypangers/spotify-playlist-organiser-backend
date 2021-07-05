package server

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const prefix = "/api"

func InitRouter() {
	r := mux.NewRouter()

	r.HandleFunc(prefix+"/getSpotifyUser", getSpotifyUserHandler)
	r.HandleFunc(prefix+"/getAuthToken", getSpotifyAccessCodeHandler)
	r.HandleFunc(prefix+"/getSpotifyUserPlaylist", getSpotifyPlaylistHandler)
	r.HandleFunc(prefix+"/getSpotifyPlaylistItemList", getSpotifyPlaylistItemListHandler)
	r.HandleFunc(prefix+"/searchItem", getSpotifySearchItemResultHandler)
	r.HandleFunc(prefix+"/createNewPlaylist", createSpotifyNewPlaylistHandler)

	http.Handle("/", r)

	http.ListenAndServe(":3030", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(r))
}
