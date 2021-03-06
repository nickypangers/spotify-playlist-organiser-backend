package server

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/nickypangers/spotifyreplaylist-backend/pkg/models"
	"github.com/nickypangers/spotifyreplaylist-backend/pkg/spotify"
)

func setHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", cors_origin)
	w.Header().Set("Content-Type", "application/json")

}

func GetSpotifyAccessTokenHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)

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
	setHeaders(w)

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

	setHeaders(w)

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

func getSpotifyPlaylistDetailHandler(w http.ResponseWriter, r *http.Request) {

	setHeaders(w)

	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)

	accessToken := r.FormValue("accessToken")
	playlistID := r.FormValue("playlistId")

	// code := r.URL.Query().Get("code")

	// log.Printf("accessToken=%v\n", accessToken)
	// log.Printf("userId=%v\n", userId)

	if len(accessToken) == 0 {
		log.Println("accessToken is empty.")
		enc.Encode("accessToken is empty")
	} else if len(playlistID) == 0 {
		log.Println("playlistID is empty.")
		enc.Encode("playlistID is empty")
	} else {
		response, _ := spotify.GetPlaylistDetail(playlistID, accessToken)

		enc.Encode(response)
	}
}

func getSpotifyPlaylistHandler(w http.ResponseWriter, r *http.Request) {

	setHeaders(w)

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

	setHeaders(w)

	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)

	playlistId := r.FormValue("playlistId")
	offset := r.FormValue("offset")
	limit := r.FormValue("limit")
	// country := r.FormValue("country")
	accessToken := r.FormValue("accessToken")

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
	setHeaders(w)

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
	setHeaders(w)

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
	setHeaders(w)

	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)

	playlistId := r.FormValue("playlistID")
	accessToken := r.FormValue("accessToken")

	response, _ := spotify.UnfollowPlaylist(playlistId, accessToken)

	enc.Encode(response)

}

func reorderPlaylistItemHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)

	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)

	playlistId := r.FormValue("playlistID")
	rangeStartStr := r.FormValue("rangeStart")
	insertBeforeStr := r.FormValue("insertBefore")
	rangeLengthStr := r.FormValue("rangeLength")
	snapshotId := r.FormValue("snapshotID")
	accessToken := r.FormValue("accessToken")

	// log.Printf("range_start=%s\ninsert_before=%s\nrange_length=%s\n", rangeStartStr, insertBeforeStr, rangeLengthStr)

	rangeStart, _ := strconv.Atoi(rangeStartStr)
	insertBefore, _ := strconv.Atoi(insertBeforeStr)
	rangeLength, _ := strconv.Atoi(rangeLengthStr)

	response, _ := spotify.ReorderPlaylistItem(rangeStart, insertBefore, rangeLength, playlistId, snapshotId, accessToken)

	enc.Encode(response)
}

func getTrackHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Content-Type", "application/json")

	enc := json.NewEncoder(w)

	enc.SetEscapeHTML(false)

	id := r.FormValue("id")
	accessToken := r.FormValue("accessToken")

	response, _ := spotify.GetTrack(id, accessToken)

	enc.Encode(response)

}

func addItemsToPlaylistHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)

	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)

	playlistId := r.FormValue("playlistId")
	position := r.FormValue("position")
	uris := r.FormValue("uris")
	accessToken := r.FormValue("accessToken")

	response, _ := spotify.AddItemsToPlaylist(playlistId, position, uris, accessToken)

	enc.Encode(response)
}

func removeItemsFromPlaylistHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)

	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)

	playlistId := r.FormValue("playlistId")
	uri := r.FormValue("uri")
	accessToken := r.FormValue("accessToken")

	response, _ := spotify.RemoveItemFromPlaylist(playlistId, uri, accessToken)

	enc.Encode(response)
}

func changePlaylistDetailHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)

	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)

	playlistId := r.FormValue("playlistId")
	playlistDetail := r.FormValue("playlistDetail")
	accessToken := r.FormValue("accessToken")

	log.Printf("playlistDetail: %s", playlistDetail)

	params, err := url.ParseQuery(playlistDetail)
	if err != nil {
		enc.Encode(models.SpotifyChangePlaylistDetailResponse{Error: struct {
			Status  int    "json:\"status\""
			Message string "json:\"message\""
		}{Status: 400, Message: "Cannot read playlistDetail"}})
		return
	}

	isPublic, _ := strconv.ParseBool(params.Get("public"))
	isCollaborative, _ := strconv.ParseBool(params.Get("collaborative"))

	jsonData := models.PlaylistDetail{Name: params.Get("name"), Public: isPublic, Collaborative: isCollaborative}

	response, _ := spotify.ChangePlaylistDetail(playlistId, accessToken, jsonData)

	enc.Encode(response)

}
