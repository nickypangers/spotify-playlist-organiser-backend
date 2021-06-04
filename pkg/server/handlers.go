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
		enc.Encode("Code is empty")
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
		log.Println("Code is empty.")
		enc.Encode("Code is empty")
	} else {
		response, _ := spotify.GetUserDetail(accessToken)

		enc.Encode(response)
	}

}
