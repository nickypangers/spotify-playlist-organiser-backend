package spotify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/nickypangers/spotifyreplaylist-backend/pkg/models"
)

func GetPlaylistItemList(playlistId, country, accessToken string) (models.SpotifyPlaylistItemList, bool) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/playlists/"+playlistId+"/tracks", nil)

	if err != nil {
		log.Println(err)
		return models.SpotifyPlaylistItemList{}, false
	}

	req.Header.Add("Authorization", "Bearer "+accessToken)

	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return models.SpotifyPlaylistItemList{}, false
	}

	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
		return models.SpotifyPlaylistItemList{}, false
	}

	var spotifyPlaylistItemList models.SpotifyPlaylistItemList

	err = json.Unmarshal(respBody, &spotifyPlaylistItemList)

	if err != nil {
		log.Println(err)
		return models.SpotifyPlaylistItemList{}, false
	}

	log.Println(spotifyPlaylistItemList)

	return spotifyPlaylistItemList, true

}

func CreateNewPlaylist(userId, playlistName, isPublic, isCollaborative, description, accessToken string) (models.SpotifyCreatePlaylistResult, bool) {

	client := &http.Client{}

	requestBody, err := json.Marshal(map[string]string{
		"name":          playlistName,
		"public":        isPublic,
		"collaborative": isCollaborative,
		"description":   description,
	})

	// data := url.Values{}

	// data.Set("name", playlistName)
	// data.Set("public", isPublic)
	// data.Set("collaborative", isCollaborative)
	// data.Set("description", description)

	// fmt.Println(data)

	// b := bytes.NewBufferString(data.Encode())

	if err != nil {
		log.Println(err)
		return models.SpotifyCreatePlaylistResult{}, false
	}

	req, err := http.NewRequest("POST", "https://api.spotify.com/v1/users/"+userId+"/playlists", bytes.NewBuffer(requestBody))

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+accessToken)

	if err != nil {
		log.Println(err)
		return models.SpotifyCreatePlaylistResult{}, false
	}

	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return models.SpotifyCreatePlaylistResult{}, false
	}

	respBody, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(respBody))

	if err != nil {
		log.Println(err)
		return models.SpotifyCreatePlaylistResult{}, false
	}

	var spotifyCreatePlaylistResult models.SpotifyCreatePlaylistResult

	err = json.Unmarshal(respBody, &spotifyCreatePlaylistResult)

	if err != nil {
		log.Println(err)
		return models.SpotifyCreatePlaylistResult{}, false
	}

	return spotifyCreatePlaylistResult, true

}
