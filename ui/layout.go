package ui

import (
	"github.com/rivo/tview"

	"github.com/evansloan/nfl-term/api"
)

// Layout represents the overall layout of the app.
// Contains all UI elements to be displayed within the window
type Layout struct {
	Outer *tview.Grid
	Inner *tview.Grid

	GameList     *GameList
	Scoreboard   *Scoreboard
	GameInfo     *GameInfo
	ScoreSummary *ScoreSummary
	TeamStats    *TeamStats
	PassStats    *PlayerStats
	RushStats    *PlayerStats
	RecvStats    *PlayerStats
}

// NewLayout creates a new Layout and arranges all
// UI elements within it.
func NewLayout() *Layout {
	l := &Layout{
		Outer:        tview.NewGrid(),
		Inner:        tview.NewGrid(),
		GameList:     NewGameList(),
		Scoreboard:   NewScoreboard(),
		GameInfo:     NewGameInfo(),
		ScoreSummary: NewScoreSummary(),
		TeamStats:    NewTeamStats(),
		PassStats:    NewPlayerStats("Passing", 0),
		RushStats:    NewPlayerStats("Rushing", 1),
		RecvStats:    NewPlayerStats("Receiving", 2),
	}

	l.Outer.SetColumns(20, 0).
		AddItem(l.GameList, 0, 0, 12, 1, 0, 5, true).
		AddItem(l.Inner, 0, 1, 12, 12, 0, 0, false)

	l.Inner.SetBorder(true).
		SetTitle("Stats")

	l.Inner.SetRows(5, 7, 8, 0).
		SetColumns(26, 40, 30, 30, 0).
		AddItem(l.Scoreboard, 0, 0, 1, 1, 0, 0, false).
		AddItem(l.GameInfo, 1, 0, 1, 1, 0, 0, false).
		AddItem(l.TeamStats, 2, 0, 1, 1, 0, 0, false).
		AddItem(l.PassStats, 0, 1, 3, 1, 0, 0, false).
		AddItem(l.RushStats, 0, 2, 3, 1, 0, 0, false).
		AddItem(l.RecvStats, 0, 3, 3, 1, 0, 0, false).
		AddItem(l.ScoreSummary, 3, 0, 20, 4, 0, 0, false)

	l.setGameList()

	return l
}

func (l *Layout) gameSelect(game *api.Game) func() {
	return func() {
		game.Update()
		l.Scoreboard.SetScores(game)
		l.GameInfo.SetInfo(game)
		l.TeamStats.SetStats(game)
		l.PassStats.SetStats(game)
		l.RushStats.SetStats(game)
		l.RecvStats.SetStats(game)
		l.ScoreSummary.SetScoreSummary(game)
	}
}

func (l *Layout) setGameList() {
	l.GameList.Clear()

	games := api.Games()
	gameSelectors := []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'g'}

	for i, game := range games {
		title := game.Home.Abbr + " vs. " + game.Away.Abbr
		callback := l.gameSelect(game)
		l.GameList.AddItem(title, "", gameSelectors[i], callback)
	}
}
