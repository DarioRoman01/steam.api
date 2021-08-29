package steam

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client interface {
	GetNewsForApp(int, int, int)
	GetGlobalAchievementPercentagesForApp()
	GetPlayerSummaries()
	GetFriendList()
	GetPlayerAchievements()
	GetUserStatsForGame()
	GetOwnedGames()
	GetRecentlyPlayedGames()
}

func NewClient(key string) Client {
	return nil
}

type SteamClient struct {
	apiKey  string
	baseUrl string
}

func (s *SteamClient) request(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

func (s *SteamClient) GetNewsForApp(appID int, count, maxLen int) {
	url := fmt.Sprintf(
		"%s/ISteamNews/GetNewsForApp/v0002/?appid=%d&count=%d&maxlength=%d&format=json",
		s.baseUrl, appID, count, maxLen,
	)

}
