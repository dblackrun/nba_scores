package main

import (
	"fmt"
	"os"

	"nba_scores/broadcast"
	"nba_scores/scoreboard"

	"github.com/jedib0t/go-pretty/table"
)

func MakeScoreboardTable(games []scoreboard.Game, gameIdBroadcasterMap map[string]string) table.Writer {
	todaysScores := table.NewWriter()
	todaysScores.SetOutputMirror(os.Stdout)
	todaysScores.AppendHeader(table.Row{"Game", "Status", "TV", "Win Prob", "Home Leader", "Away Leader"})
	for _, game := range games {
		gameValue := fmt.Sprintf("%s (%d) @ %s (%d)", game.AwayTeam.TeamTricode, game.AwayTeam.Score, game.HomeTeam.TeamTricode, game.HomeTeam.Score)
		hl := game.GameLeaders.HomeLeaders
		homeLeader := fmt.Sprintf("%s - %dPTS %dREB %dAST", hl.Name, hl.Points, hl.Rebounds, hl.Assists)
		al := game.GameLeaders.AwayLeaders
		awayLeader := fmt.Sprintf("%s - %dPTS %dREB %dAST", al.Name, al.Points, al.Rebounds, al.Assists)
		winProb := fmt.Sprintf("%s %f", game.PbOdds.Team, game.PbOdds.Odds)
		tv, found := gameIdBroadcasterMap[game.GameID]
		if !found {
			tv = ""
		}
		todaysScores.AppendRow([]interface{}{gameValue, game.GameStatusText, tv, winProb, homeLeader, awayLeader})
	}
	return todaysScores
}

func main() {
	broadcastResponse, err := broadcast.GetBroadcastResponseData()
	if err != nil {
		fmt.Println(err.Error())
	}
	gameIdBroadcasterMap := make(map[string]string)
	for _, game := range broadcastResponse.ResultSets[1].CompleteGameList {
		gameIdBroadcasterMap[game.GameID] = game.BroadcasterName
	}

	scoreboardResponse, err := scoreboard.GetScoreboardResponseData()
	if err != nil {
		fmt.Println(err.Error())
	}

	todaysGames := MakeScoreboardTable(scoreboardResponse.Scoreboard.Games, gameIdBroadcasterMap)
	todaysGames.Render()
}
