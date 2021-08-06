package spotify

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/nickypangers/spotifyreplaylist-backend/pkg/models"
)

func GetPlaylistItemList(playlistId, offset, limit, accessToken string) (models.SpotifyPlaylistItemListResponse, bool) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/playlists/"+playlistId+"/tracks", nil)

	if err != nil {
		log.Println(err)
		return models.SpotifyPlaylistItemListResponse{}, false
	}

	q := url.Values{}
	q.Add("offset", offset)
	q.Add("limit", limit)

	req.URL.RawQuery = q.Encode()

	req.Header.Add("Authorization", "Bearer "+accessToken)

	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return models.SpotifyPlaylistItemListResponse{}, false
	}

	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
		return models.SpotifyPlaylistItemListResponse{}, false
	}

	var spotifyPlaylistItemListResponse models.SpotifyPlaylistItemListResponse

	err = json.Unmarshal(respBody, &spotifyPlaylistItemListResponse)

	if err != nil {
		log.Println(err)
		return models.SpotifyPlaylistItemListResponse{}, false
	}

	log.Printf("%s get playlist item list: %s", accessToken, playlistId)

	return spotifyPlaylistItemListResponse, true

}

func CreateNewPlaylist(userId, playlistName, isPublic, isCollaborative, description, accessToken string) (models.SpotifyCreatePlaylistResponse, bool) {

	client := &http.Client{}

	requestBody, err := json.Marshal(map[string]string{
		"name":          playlistName,
		"public":        isPublic,
		"collaborative": isCollaborative,
		"description":   description,
	})

	if err != nil {
		log.Println(err)
		return models.SpotifyCreatePlaylistResponse{}, false
	}

	req, err := http.NewRequest("POST", "https://api.spotify.com/v1/users/"+userId+"/playlists", bytes.NewBuffer(requestBody))

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+accessToken)

	if err != nil {
		log.Println(err)
		return models.SpotifyCreatePlaylistResponse{}, false
	}

	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return models.SpotifyCreatePlaylistResponse{}, false
	}

	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
		return models.SpotifyCreatePlaylistResponse{}, false
	}

	var spotifyCreatePlaylistResponse models.SpotifyCreatePlaylistResponse

	err = json.Unmarshal(respBody, &spotifyCreatePlaylistResponse)

	if err != nil {
		log.Println(err)
		return models.SpotifyCreatePlaylistResponse{}, false
	}

	log.Printf("%s created new playlist: %s", accessToken, playlistName)

	return spotifyCreatePlaylistResponse, true

}

func ReorderPlaylistItem(rangeStart, insertBefore, rangeLength int, playlistId, snapshotId, accessToken string) (models.SpotifyReorderPlaylistItemResponse, bool) {

	client := &http.Client{}

	requestBody, err := json.Marshal(map[string]int{
		"range_start":   rangeStart,
		"insert_before": insertBefore,
		"range_length":  rangeLength,
	})

	if err != nil {
		return models.SpotifyReorderPlaylistItemResponse{}, false
	}

	req, err := http.NewRequest("PUT", "https://api.spotify.com/v1/playlists/3K0P6CWMe5Iy663QVOpF8E/tracks", bytes.NewBuffer(requestBody))

	req.Header.Add("Content-Type", "application/json")

	req.Header.Add("Authorization", "Bearer "+accessToken)

	if err != nil {
		log.Println(err)
		return models.SpotifyReorderPlaylistItemResponse{}, false
	}

	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return models.SpotifyReorderPlaylistItemResponse{}, false
	}

	respBody, err := ioutil.ReadAll(resp.Body)

	log.Println(string(respBody))

	if err != nil {
		log.Println(err)
		return models.SpotifyReorderPlaylistItemResponse{}, false
	}

	var spotifyReorderPlaylistItemResponse models.SpotifyReorderPlaylistItemResponse

	err = json.Unmarshal(respBody, &spotifyReorderPlaylistItemResponse)

	if err != nil {
		log.Println(err)
		return models.SpotifyReorderPlaylistItemResponse{}, false
	}

	log.Printf("%s reordered playlist item on playlist: %s", accessToken, playlistId)

	return spotifyReorderPlaylistItemResponse, true

}

func AddItemsToPlaylist(playlistId, position, uris, accessToken string) (models.SpotifyAddItemToPlaylistResponse, bool) {

	client := &http.Client{}

	req, err := http.NewRequest("PUT", "https://api.spotify.com/v1/playlists/3K0P6CWMe5Iy663QVOpF8E/tracks", nil)

	qs := req.URL.Query()

	qs.Add("position", position)
	qs.Add("uris", uris)

	req.URL.RawQuery = qs.Encode()

	req.Header.Add("Content-Type", "application/json")

	req.Header.Add("Authorization", "Bearer "+accessToken)

	if err != nil {
		log.Println(err)
		return models.SpotifyAddItemToPlaylistResponse{}, false
	}

	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return models.SpotifyAddItemToPlaylistResponse{}, false
	}

	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
		return models.SpotifyAddItemToPlaylistResponse{}, false
	}

	var spotifyAddItemToPlaylist models.SpotifyAddItemToPlaylistResponse

	err = json.Unmarshal(respBody, &spotifyAddItemToPlaylist)

	if err != nil {
		log.Println(err)
		return models.SpotifyAddItemToPlaylistResponse{}, false
	}

	log.Printf("%s added playlist item: %s to playlist: %s", accessToken, uris, playlistId)

	return spotifyAddItemToPlaylist, true

}
