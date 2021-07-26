package spotify

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/nickypangers/spotifyreplaylist-backend/pkg/models"
)

func GetUserDetail(accessToken string) (models.SpotifyProfileResponse, bool) {

	// resp, err := http.Get("https://api.spotify.com/v1/me")

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/me", nil)

	if err != nil {
		log.Println(err)
		return models.SpotifyProfileResponse{}, false
	}

	req.Header.Add("Authorization", "Bearer "+accessToken)

	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return models.SpotifyProfileResponse{}, false
	}

	respBody, err := ioutil.ReadAll(resp.Body)

	// log.Println(string(respBody))

	if err != nil {
		log.Println(err)
		return models.SpotifyProfileResponse{}, false
	}

	var spotifyProfileResponse models.SpotifyProfileResponse

	err = json.Unmarshal(respBody, &spotifyProfileResponse)

	if err != nil {
		log.Println(err)
		return models.SpotifyProfileResponse{}, false
	}

	// log.Println(spotifyProfileResponse)

	return spotifyProfileResponse, true

}

func GetUserPlaylists(userId, accessToken string) (models.SpotifyUserPlaylistResponse, bool) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/users/"+userId+"/playlists", nil)

	if err != nil {
		log.Println(err)
		return models.SpotifyUserPlaylistResponse{}, false
	}

	req.Header.Add("Authorization", "Bearer "+accessToken)

	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return models.SpotifyUserPlaylistResponse{}, false
	}

	respBody, err := ioutil.ReadAll(resp.Body)

	// log.Println(string(respBody))

	if err != nil {
		log.Println(err)
		return models.SpotifyUserPlaylistResponse{}, false
	}

	var spotifyUserPlaylistResponse models.SpotifyUserPlaylistResponse

	err = json.Unmarshal(respBody, &spotifyUserPlaylistResponse)

	if err != nil {
		log.Println(err)
		return models.SpotifyUserPlaylistResponse{}, false
	}

	// log.Println(spotifyUserPlaylistResponse)

	return spotifyUserPlaylistResponse, true
}
