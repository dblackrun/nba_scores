package scoreboard

import "time"

type ScoreboardBody struct {
	Meta struct {
		Version int    `json:"version"`
		Request string `json:"request"`
		Time    string `json:"time"`
		Code    int    `json:"code"`
	} `json:"meta"`
	Scoreboard struct {
		GameDate   string `json:"gameDate"`
		LeagueID   string `json:"leagueId"`
		LeagueName string `json:"leagueName"`
		Games      []Game `json:"games"`
	} `json:"scoreboard"`
}

type Game struct {
	GameID            string    `json:"gameId"`
	GameCode          string    `json:"gameCode"`
	GameStatus        int       `json:"gameStatus"`
	GameStatusText    string    `json:"gameStatusText"`
	Period            int       `json:"period"`
	GameClock         string    `json:"gameClock"`
	GameTimeUTC       time.Time `json:"gameTimeUTC"`
	GameEt            string    `json:"gameEt"`
	RegulationPeriods int       `json:"regulationPeriods"`
	IfNecessary       bool      `json:"ifNecessary"`
	SeriesGameNumber  string    `json:"seriesGameNumber"`
	SeriesText        string    `json:"seriesText"`
	HomeTeam          Team      `json:"homeTeam"`
	AwayTeam          Team      `json:"awayTeam"`
	GameLeaders       struct {
		HomeLeaders TeamLeaders `json:"homeLeaders"`
		AwayLeaders TeamLeaders `json:"awayLeaders"`
	} `json:"gameLeaders"`
	PbOdds struct {
		Team      string  `json:"team"`
		Odds      float64 `json:"odds"`
		Suspended int     `json:"suspended"`
	} `json:"pbOdds"`
}

type Team struct {
	TeamID            int         `json:"teamId"`
	TeamName          string      `json:"teamName"`
	TeamCity          string      `json:"teamCity"`
	TeamTricode       string      `json:"teamTricode"`
	Wins              int         `json:"wins"`
	Losses            int         `json:"losses"`
	Score             int         `json:"score"`
	Seed              interface{} `json:"seed"`
	InBonus           string      `json:"inBonus"`
	TimeoutsRemaining int         `json:"timeoutsRemaining"`
	Periods           []struct {
		Period     int    `json:"period"`
		PeriodType string `json:"periodType"`
		Score      int    `json:"score"`
	} `json:"periods"`
}

type TeamLeaders struct {
	PersonID    int    `json:"personId"`
	Name        string `json:"name"`
	JerseyNum   string `json:"jerseyNum"`
	Position    string `json:"position"`
	TeamTricode string `json:"teamTricode"`
	PlayerSlug  string `json:"playerSlug"`
	Points      int    `json:"points"`
	Rebounds    int    `json:"rebounds"`
	Assists     int    `json:"assists"`
}
