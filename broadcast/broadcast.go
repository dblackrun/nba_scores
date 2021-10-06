package broadcast

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetBroadcastResponseData() (BroadcastResponseBody, error) {
	var result BroadcastResponseBody
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://stats.nba.com/stats/internationalbroadcasterschedule", nil)
	if err != nil {
		return result, err
	}

	q := req.URL.Query()
	q.Add("LeagueID", "00")
	q.Add("Season", "2021")
	q.Add("RegionID", "2")
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; rv:78.0) Gecko/20100101 Firefox/78.0")
	req.Header.Set("Referer", "http://global.nba.com/broadcaster-schedule/")
	req.Header.Set("Origin", "http://global.nba.com")

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
