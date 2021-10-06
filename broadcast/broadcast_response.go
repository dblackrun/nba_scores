package broadcast

type BroadcastResponseBody struct {
	Resource   string `json:"resource"`
	Parameters struct {
		LeagueID string      `json:"LeagueID"`
		Season   string      `json:"Season"`
		RegionID int         `json:"RegionID"`
		Date     interface{} `json:"Date"`
		Est      interface{} `json:"EST"`
	} `json:"parameters"`
	ResultSets []struct {
		CompleteGameList []Game `json:"CompleteGameList"`
	} `json:"resultSets"`
}

type Game struct {
	GameID            string `json:"gameID"`
	VtCity            string `json:"vtCity"`
	VtNickName        string `json:"vtNickName"`
	VtShortName       string `json:"vtShortName"`
	VtAbbreviation    string `json:"vtAbbreviation"`
	HtCity            string `json:"htCity"`
	HtNickName        string `json:"htNickName"`
	HtShortName       string `json:"htShortName"`
	HtAbbreviation    string `json:"htAbbreviation"`
	Date              string `json:"date"`
	Time              string `json:"time"`
	Day               string `json:"day"`
	BroadcastID       string `json:"broadcastID"`
	BroadcasterName   string `json:"broadcasterName"`
	TapeDelayComments string `json:"tapeDelayComments"`
}
