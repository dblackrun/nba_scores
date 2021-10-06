package scoreboard

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetScoreboardResponseData() (ScoreboardBody, error) {
	var result ScoreboardBody
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://nba-prod-us-east-1-mediaops-stats.s3.amazonaws.com/NBA/liveData/scoreboard/todaysScoreboard_00.json", nil)
	if err != nil {
		return result, err
	}

	// req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	// req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; rv:78.0) Gecko/20100101 Firefox/78.0")
	// req.Header.Set("Referer", "http://global.nba.com/broadcaster-schedule/")
	// req.Header.Set("Origin", "http://global.nba.com")

	response, err := client.Do(req)
	if err != nil {
		return result, err
	}
	defer func() { _ = response.Body.Close() }()
	if response.StatusCode != 200 {
		return result, err
	}
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(contents, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
