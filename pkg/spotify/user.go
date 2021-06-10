package spotify

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/nickypangers/spotifyreplaylist-backend/pkg/models"
)

func GetUserDetail(accessToken string) (models.SpotifyProfile, bool) {

	// resp, err := http.Get("https://api.spotify.com/v1/me")

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/me", nil)

	if err != nil {
		log.Println(err)
		return models.SpotifyProfile{}, false
	}

	req.Header.Add("Authorization", "Bearer "+accessToken)

	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return models.SpotifyProfile{}, false
	}

	respBody, err := ioutil.ReadAll(resp.Body)

	// log.Println(string(respBody))

	if err != nil {
		log.Println(err)
		return models.SpotifyProfile{}, false
	}

	var spotifyProfile models.SpotifyProfile

	err = json.Unmarshal(respBody, &spotifyProfile)

	if err != nil {
		log.Println(err)
		return models.SpotifyProfile{}, false
	}

	// log.Println(spotifyProfile)

	return spotifyProfile, true

}

func GetUserPlaylists(userId, accessToken string) (models.SpotifyUserPlaylist, bool) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/users/"+userId+"/playlists", nil)

	if err != nil {
		log.Println(err)
		return models.SpotifyUserPlaylist{}, false
	}

	req.Header.Add("Authorization", "Bearer "+accessToken)

	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return models.SpotifyUserPlaylist{}, false
	}

	respBody, err := ioutil.ReadAll(resp.Body)

	// log.Println(string(respBody))

	if err != nil {
		log.Println(err)
		return models.SpotifyUserPlaylist{}, false
	}

	var spotifyUserPlaylist models.SpotifyUserPlaylist

	err = json.Unmarshal(respBody, &spotifyUserPlaylist)

	if err != nil {
		log.Println(err)
		return models.SpotifyUserPlaylist{}, false
	}

	// log.Println(spotifyUserPlaylist)

	return spotifyUserPlaylist, true
}
