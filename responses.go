package steam

// Friends related responses
type FriendList struct {
	Friends []*Friend `json:"friends"`
}

type FriendResponse struct {
	FriendList FriendList `json:"friendslist"`
}

type Friend struct {
	Steamid      string `json:"steamid"`
	Relationship string `json:"relationship"`
	FriendSince  int    `json:"friend_since"`
}

// Game stats related responses
type PLayerStatsResponse struct {
	PlayerStats PlayerStats
}

type PlayerStats struct {
	SteamId  string
	GameName string
	Stats    []*Stats
}

type Stats struct {
	Name  string
	Value int
}
