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

// User info related responses
type UserInfoResponse struct {
	Response struct {
		Players Player `json:"players"`
	} `json:"response"`
}

type Player struct {
	SteamID                  string `json:"steamid"`
	Communityvisibilitystate int    `json:"communityvisibilitystate"`
	Profilestate             int    `json:"profilestate"`
	PersonName               string `json:"personaname"`
	ProfileUrl               string `json:"profileurl"`
	Avatar                   string `json:"avatar"`
	AvatarMedium             string `json:"avatarmedium"`
	AvatarFull               string `json:"avatarfull"`
	AvatarHash               string `json:"avatarhash"`
	PersonState              int    `json:"personastate"`
	RealName                 string `json:"realname"`
	Primaryclanid            string `json:"primaryclanid"`
	TimeCreated              int    `json:"timecreated"`
	PersonStateFlag          int    `json:"personastateflags"`
	LocCounttryCode          string `json:"loccountrycode"`
	LocStateCode             string `json:"locstatecode"`
	LocCityID                int    `json:"loccityid"`
}

type GamesResponse struct {
	Response struct {
		TotalCount int    `json:"total_count"`
		Games      []Game `json:"games"`
	} `json:"response"`
}

type Game struct {
	AppID                  int    `json:"appid"`
	Name                   string `json:"name"`
	PlayTime2Weeks         int    `json:"playtime_2weeks"`
	PlayTimeForever        int    `json:"playtime_forever"`
	ImgIconUrl             string `json:"img_icon_url"`
	ImgLogoUrl             string `json:"img_logo_url"`
	PlayTimeWIndowsForever int    `json:"playtime_windows_forever"`
	PlayTimeMacForever     int    `json:"playtime_mac_forever"`
	PlayTimeLinuxForever   int    `json:"playtime_linux_forever"`
}
