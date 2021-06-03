package spotify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/nickypangers/spotifyreplaylist-backend/pkg/models"
)

func GetUserDetail(accessCode string) (models.SpotifyProfile, bool) {

	// resp, err := http.Get("https://api.spotify.com/v1/me")

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/me", nil)

	if err != nil {
		log.Println(err)
		return models.SpotifyProfile{}, false
	}

	req.Header.Add("Authorization", "Bearer "+accessCode)

	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return models.SpotifyProfile{}, false
	}

	respBody, err := ioutil.ReadAll(resp.Body)

	// fmt.Println(string(respBody))

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

	fmt.Println(spotifyProfile)

	return spotifyProfile, true

}
