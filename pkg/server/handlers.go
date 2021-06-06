package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/nickypangers/spotifyreplaylist-backend/pkg/spotify"
)

func GetSpotifyAccessCodeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Content-Type", "application/json")

	enc := json.NewEncoder(w)

	enc.SetEscapeHTML(false)

	code := r.FormValue("code")

	// code := r.URL.Query().Get("code")

	fmt.Printf("code=%v\n", code)

	if len(code) == 0 {
		log.Println("Code is empty.")
		enc.Encode("accessToken is empty")
	} else {
		accessCode, status := spotify.GetSpotifyAccessCode(code)

		if !status {
			log.Println("Unable to get spotify user.")
		} else {
			// response, _ := spotify.GetUserDetail(accessCode)

			enc.Encode(accessCode)
		}
	}
}

func getSpotifyUserHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Content-Type", "application/json")

	enc := json.NewEncoder(w)

	enc.SetEscapeHTML(false)

	accessToken := r.FormValue("accessToken")

	// code := r.URL.Query().Get("code")

	fmt.Printf("accessToken=%v\n", accessToken)

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

	fmt.Printf("accessToken=%v\n", accessToken)
	fmt.Printf("userId=%v\n", userId)

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
	country := r.FormValue("country")
	accessToken := r.FormValue("accessToken")

	// code := r.URL.Query().Get("code")

	fmt.Printf("accessToken=%v\n", accessToken)
	fmt.Printf("playlistId=%v\n", playlistId)
	fmt.Printf("country=%v\n", country)

	if len(accessToken) == 0 {
		log.Println("accessToken is empty.")
		enc.Encode("accessToken is empty")
	} else if len(playlistId) == 0 {
		log.Println("playlistId is empty.")
		enc.Encode("playlistId is empty")
	} else if len(country) == 0 {
		log.Println("country is empty.")
		enc.Encode("country is empty")
	} else {
		response, _ := spotify.GetPlaylistItemList(playlistId, country, accessToken)

		enc.Encode(response)
	}
}
