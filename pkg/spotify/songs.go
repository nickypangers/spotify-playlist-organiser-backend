package spotify

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/nickypangers/spotifyreplaylist-backend/pkg/models"
)

func SearchItem(q, t, accessToken string) (models.SpotifySearchItemResponse, bool) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/search", nil)

	if err != nil {
		log.Println(err)
		return models.SpotifySearchItemResponse{}, false
	}

	q = strings.ReplaceAll(q, " ", "+")

	qs := req.URL.Query()

	qs.Add("q", "\""+q+"\"")

	qs.Add("type", t)

	qs.Add("market", "from_token")

	req.URL.RawQuery = qs.Encode()

	req.Header.Add("Authorization", "Bearer "+accessToken)

	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return models.SpotifySearchItemResponse{}, false
	}

	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
		return models.SpotifySearchItemResponse{}, false
	}

	// log.Println(string(respBody))

	var spotifySearchItemResponse models.SpotifySearchItemResponse

	err = json.Unmarshal(respBody, &spotifySearchItemResponse)

	if err != nil {
		log.Println(err)
		return models.SpotifySearchItemResponse{}, false
	}

	log.Printf("%s query song: q = %s, t = %s", accessToken, q, t)

	return spotifySearchItemResponse, true

}

func GetTrack(id, accessToken string) (models.SpotifyTrackResponse, bool) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/tracks/"+id, nil)
	if err != nil {
		log.Println(err)
		return models.SpotifyTrackResponse{}, false
	}

	qs := req.URL.Query()

	qs.Add("market", "from_token")

	req.URL.RawQuery = qs.Encode()

	req.Header.Add("Authorization", "Bearer "+accessToken)

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return models.SpotifyTrackResponse{}, false
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return models.SpotifyTrackResponse{}, false
	}

	var spotifyTrackResponse models.SpotifyTrackResponse
	err = json.Unmarshal(respBody, &spotifyTrackResponse)
	if err != nil {
		log.Println(err)
		return models.SpotifyTrackResponse{}, false
	}

	log.Printf("%s get track: %s", accessToken, id)

	return spotifyTrackResponse, true

}
