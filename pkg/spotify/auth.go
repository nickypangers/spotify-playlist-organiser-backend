package spotify

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/nickypangers/spotifyreplaylist-backend/pkg/credential"
	"github.com/nickypangers/spotifyreplaylist-backend/pkg/models"
)

func getAuthHeader(prefix string) string {
	data := credential.ClientID + ":" + credential.ClientSecret

	log.Println(data)

	sEnc := base64.StdEncoding.EncodeToString([]byte(data))

	log.Println(sEnc)

	return prefix + " " + sEnc
}

func GetSpotifyAccessCode(grantType, code string) (models.AccessTokenResponse, bool) {

	authHeader := getAuthHeader("Basic")

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", nil)
	if err != nil {
		log.Fatal(err)
	}

	// set POST param
	q := url.Values{}
	q.Add("grant_type", grantType)
	q.Add("code", code)
	q.Add("redirect_uri", credential.RedirectUrl)

	// set POST header
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", authHeader)

	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return models.AccessTokenResponse{}, false
	}

	d, _ := ioutil.ReadAll(resp.Body)

	var accessTokenResponse models.AccessTokenResponse

	err = json.Unmarshal(d, &accessTokenResponse)

	if err != nil {
		log.Printf("error decoding response: %v", err)
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
		}
		log.Printf("response: %q", d)
		return models.AccessTokenResponse{}, false
	}

	return accessTokenResponse, true

}
