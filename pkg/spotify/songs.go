package spotify

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/nickypangers/spotifyreplaylist-backend/pkg/models"
)

func SearchItem(q, t, accessToken string) (models.SpotifySearchItemResult, bool) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/search", nil)

	if err != nil {
		log.Println(err)
		return models.SpotifySearchItemResult{}, false
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
		return models.SpotifySearchItemResult{}, false
	}

	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
		return models.SpotifySearchItemResult{}, false
	}

	// log.Println(string(respBody))

	var spotifySearchItemResult models.SpotifySearchItemResult

	err = json.Unmarshal(respBody, &spotifySearchItemResult)

	if err != nil {
		log.Println(err)
		return models.SpotifySearchItemResult{}, false
	}

	// log.Println(spotifySearchItemResult)

	return spotifySearchItemResult, true

}
