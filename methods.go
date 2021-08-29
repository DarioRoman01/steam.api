package steam

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Client is an interface to implements all the steam web api methods
type Client interface {
	GetNewsForApp(int, int, int)
	GetGlobalAchievementPercentagesForApp(int)
	GetPlayerSummaries(int)
	GetFriendList(int)
	GetPlayerAchievements()
	GetUserStatsForGame()
	GetOwnedGames()
	GetRecentlyPlayedGames()
}

// create a new client instance
func NewClient(key string) Client {
	return nil
}

// represents a client for the steam api
type SteamClient struct {
	apiKey  string // represents the api key
	baseUrl string // holds the base url of the api
}

// utility function to
func (s *SteamClient) request(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

func (s *SteamClient) GetNewsForApp(appID int, count, maxLen int) (interface{}, error) {
	url := fmt.Sprintf(
		"%s/ISteamNews/GetNewsForApp/v0002/?appid=%d&count=%d&maxlength=%d&format=json",
		s.baseUrl, appID, count, maxLen,
	)

	data, err := s.request(url)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *SteamClient) GetGlobalAchievementPercentagesForApp(appID int) (interface{}, error) {
	url := fmt.Sprintf("%s/GetGlobalAchievementPercentagesForApp/v0002/?gameid=%d", s.baseUrl, appID)
	data, err := s.request(url)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *SteamClient) GetPlayerSummaries(steamID int) (interface{}, error) {
	url := fmt.Sprintf("%s/GetPlayerSummaries/v0002/?key=%s&steamids=%d", s.baseUrl, s.apiKey, steamID)
	data, err := s.request(url)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *SteamClient) GetFriendList(steamID int) (interface{}, error) {
	url := fmt.Sprintf("%s/GetFriendList/v0001/?key=%s&steamid=%d&relationship=friend", s.baseUrl, s.apiKey, steamID)
	data, err := makeRequest(url)
	if err != nil {
		return nil, err
	}

	return data, nil
}
