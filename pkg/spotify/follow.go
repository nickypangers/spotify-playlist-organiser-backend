package spotify

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/nickypangers/spotifyreplaylist-backend/pkg/models"
)

func UnfollowPlaylist(playlistId, accessToken string) (models.SpotifyUnfollowPlaylistResult, bool) {

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", "https://api.spotify.com/v1/playlists/"+playlistId+"/followers", nil)

	if err != nil {
		log.Println(err)
		return models.SpotifyUnfollowPlaylistResult{Success: false}, false
	}

	req.Header.Add("Authorization", "Bearer "+accessToken)

	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return models.SpotifyUnfollowPlaylistResult{Success: false}, false
	}

	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
		return models.SpotifyUnfollowPlaylistResult{Success: false}, false
	}

	log.Println(string(respBody))

	return models.SpotifyUnfollowPlaylistResult{Success: true}, true

}
