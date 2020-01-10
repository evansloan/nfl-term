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
	PlayerStats  *StatPages
	KeyBinds     *tview.TextView
}

// NewLayout creates a new Layout and arranges all
// UI elements within it.
func NewLayout(games []*api.Game) *Layout {
	l := &Layout{
		Outer:        tview.NewGrid(),
		Inner:        tview.NewGrid(),
		GameList:     NewGameList(),
		Scoreboard:   NewScoreboard(),
		GameInfo:     NewGameInfo(),
		ScoreSummary: NewScoreSummary(),
		TeamStats:    NewTeamStats(),
		PlayerStats:  NewStatPages(),
		KeyBinds:     tview.NewTextView(),
	}

	l.KeyBinds.SetText("Ctrl+D - Switch to defense stats | Ctrl+O - Switch to offense stats | TAB - Change stat category | ESC - Back to game list")
	l.KeyBinds.SetTextAlign(1)

	l.Outer.SetRows(0, 1).
		SetColumns(20, 0).
		AddItem(l.GameList, 0, 0, 1, 1, 0, 5, true).
		AddItem(l.Inner, 0, 1, 1, 12, 0, 0, false).
		AddItem(l.KeyBinds, 1, 0, 1, 12, 0, 0, false)

	l.Inner.SetBorder(true).
		SetTitle("Stats")

	l.Inner.SetRows(5, 7, 8, 0).
		SetColumns(26, 0).
		AddItem(l.Scoreboard, 0, 0, 1, 1, 0, 0, false).
		AddItem(l.GameInfo, 1, 0, 1, 1, 0, 0, false).
		AddItem(l.TeamStats, 2, 0, 1, 1, 0, 0, false).
		AddItem(l.PlayerStats, 0, 1, 3, 4, 0, 0, true).
		AddItem(l.ScoreSummary, 3, 0, 20, 5, 0, 0, true)

	l.setGameList(games)

	return l
}

func (l *Layout) setGameList(games []*api.Game) {
	l.GameList.Clear()

	gameSelectors := []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'g'}

	for i, game := range games {
		title := game.Home.Abbr + " vs. " + game.Away.Abbr
		l.GameList.AddItem(title, "", gameSelectors[i], l.gameSelect(game))
	}

	l.gameSelect(games[0])()
}

func (l *Layout) gameSelect(game *api.Game) func() {
	return func() {
		game.Update()
		l.Scoreboard.SetScores(game)
		l.GameInfo.SetInfo(game)
		l.TeamStats.SetStats(game)
		l.PlayerStats.SetStats(game)
		l.ScoreSummary.SetScoreSummary(game)
	}
}
