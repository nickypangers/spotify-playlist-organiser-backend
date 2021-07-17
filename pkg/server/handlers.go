package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/nickypangers/spotifyreplaylist-backend/pkg/spotify"
)

func GetSpotifyAccessTokenHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Content-Type", "application/json")

	enc := json.NewEncoder(w)

	enc.SetEscapeHTML(false)

	code := r.FormValue("code")
	grantType := r.FormValue("grantType")

	if len(code) == 0 {
		log.Println("code is empty.")
		enc.Encode("code is empty")
	} else if len(grantType) == 0 {
		log.Println("grantType is empty")
		enc.Encode("grantType is empty")
	} else {
		accessCode, status := spotify.GetSpotifyAccessToken(grantType, code)

		if !status {
			log.Println("Unable to get spotify user.")
		} else {

			enc.Encode(accessCode)
		}
	}
}

func getRefreshedAccessTokenHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Content-Type", "application/json")

	enc := json.NewEncoder(w)

	enc.SetEscapeHTML(false)

	refreshToken := r.FormValue("refreshToken")

	if len(refreshToken) == 0 {
		log.Println("refreshToken is empty")
		enc.Encode("refreshToken is empty")
	} else {
		response, _ := spotify.GetRefreshedAccessToken(refreshToken)

		enc.Encode(response)
	}

}

func getSpotifyUserHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Content-Type", "application/json")

	enc := json.NewEncoder(w)

	enc.SetEscapeHTML(false)

	accessToken := r.FormValue("accessToken")

	// code := r.URL.Query().Get("code")

	// log.Printf("accessToken=%v\n", accessToken)

	if len(accessToken) == 0 {
		log.Println("accessToken is empty.")
		enc.Encode("accessToken is empty")
	} else {
		response, _ := spotify.GetUserDetail(accessToken)

		enc.Encode(response)
	}

}

func getSpotifyPlaylistHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Content-Type", "application/json")

	enc := json.NewEncoder(w)

	enc.SetEscapeHTML(false)

	accessToken := r.FormValue("accessToken")
	userId := r.FormValue("userId")

	// code := r.URL.Query().Get("code")

	// log.Printf("accessToken=%v\n", accessToken)
	// log.Printf("userId=%v\n", userId)

	if len(accessToken) == 0 {
		log.Println("accessToken is empty.")
		enc.Encode("accessToken is empty")
	} else if len(userId) == 0 {
		log.Println("userId is empty.")
		enc.Encode("userId is empty")
	} else {
		response, _ := spotify.GetUserPlaylists(userId, accessToken)

		enc.Encode(response)
	}
}

func getSpotifyPlaylistItemListHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Content-Type", "application/json")

	enc := json.NewEncoder(w)

	enc.SetEscapeHTML(false)

	playlistId := r.FormValue("playlistId")
	offset := r.FormValue("offset")
	limit := r.FormValue("limit")
	// country := r.FormValue("country")
	accessToken := r.FormValue("accessToken")

	// code := r.URL.Query().Get("code")

	// log.Printf("accessToken=%v\n", accessToken)
	// log.Printf("playlistId=%v\n", playlistId)
	// log.Printf("country=%v\n", country)

	if len(accessToken) == 0 {
		log.Println("accessToken is empty.")
		enc.Encode("accessToken is empty")
	} else if len(playlistId) == 0 {
		log.Println("playlistId is empty.")
		enc.Encode("playlistId is empty")
	} else {
		response, _ := spotify.GetPlaylistItemList(playlistId, offset, limit, accessToken)

		enc.Encode(response)
	}
}

func getSpotifySearchItemResultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Content-Type", "application/json")

	enc := json.NewEncoder(w)

	enc.SetEscapeHTML(false)

	q := r.FormValue("q")
	t := r.FormValue("t")
	accessToken := r.FormValue("accessToken")

	if len(q) == 0 {
		log.Println("q is empty")
		enc.Encode("q is empty")
	} else if len(t) == 0 {
		log.Println("t is empty")
		enc.Encode("t is empty")
	} else if len(accessToken) == 0 {
		log.Println("accessToken is empty")
		enc.Encode("accessToken is empty")
	} else {
		response, _ := spotify.SearchItem(q, t, accessToken)

		enc.Encode(response)
	}

}

func createSpotifyNewPlaylistHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Content-Type", "application/json")

	enc := json.NewEncoder(w)

	enc.SetEscapeHTML(false)

	userId := r.FormValue("userID")
	playlistName := r.FormValue("playlistName")
	isPublic := r.FormValue("isPublic")
	isCollaborative := r.FormValue("isCollaborative")
	description := r.FormValue("description")
	accessToken := r.FormValue("accessToken")

	response, _ := spotify.CreateNewPlaylist(userId, playlistName, isPublic, isCollaborative, description, accessToken)

	enc.Encode(response)
}

func unfollowPlaylistHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Content-Type", "application/json")

	enc := json.NewEncoder(w)

	enc.SetEscapeHTML(false)

	playlistId := r.FormValue("playlistID")
	accessToken := r.FormValue("accessToken")

	response, _ := spotify.UnfollowPlaylist(playlistId, accessToken)

	enc.Encode(response)

}
