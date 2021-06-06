package spotify

import (
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

	fmt.Println(spotifyPlaylistItemList)

	return spotifyPlaylistItemList, true

}
