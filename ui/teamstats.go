package ui

import (
	"github.com/evansloan/nfl-term/api"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

// TeamStats represents a UI element that displays
// basic team stats for an NFL game.
//
// Inherits from DefaultTable
type TeamStats struct {
	*DefaultTable
}

// NewTeamStats creates a new TeamStats element.
func NewTeamStats() *TeamStats {
	return &TeamStats{
		DefaultTable: NewDefaultTable("Team Stats"),
	}
}

// SetStats populates a TeamStats element with stats
// from a specific NFL game.
func (t *TeamStats) SetStats(game *api.Game) {
	t.Clear()

	t.SetCustomCell(0, 0, game.Home.Abbr, teamColors[game.Home.Abbr], tview.AlignCenter).
		SetCustomCell(0, 2, game.Away.Abbr, teamColors[game.Away.Abbr], tview.AlignCenter)

	headerColor := tcell.ColorLightGray
	t.SetCustomCell(1, 1, "Pass", headerColor, tview.AlignCenter).
		SetCustomCell(2, 1, "Rush", headerColor, tview.AlignCenter).
		SetCustomCell(3, 1, "Total", headerColor, tview.AlignCenter).
		SetCustomCell(4, 1, "Time of Pos", headerColor, tview.AlignCenter).
		SetCustomCell(5, 1, "Timeouts", headerColor, tview.AlignCenter)

	t.SetCustomIntCell(1, 0, game.Home.Stats.Team.PassYards, tcell.ColorWhite, tview.AlignCenter).
		SetCustomIntCell(1, 2, game.Away.Stats.Team.PassYards, tcell.ColorWhite, tview.AlignCenter).
		SetCustomIntCell(2, 0, game.Home.Stats.Team.RushYards, tcell.ColorWhite, tview.AlignCenter).
		SetCustomIntCell(2, 2, game.Away.Stats.Team.RushYards, tcell.ColorWhite, tview.AlignCenter).
		SetCustomIntCell(3, 0, game.Home.Stats.Team.TotalYards, tcell.ColorWhite, tview.AlignCenter).
		SetCustomIntCell(3, 2, game.Away.Stats.Team.TotalYards, tcell.ColorWhite, tview.AlignCenter).
		SetCustomCell(4, 0, game.Home.Stats.Team.Top, tcell.ColorWhite, tview.AlignCenter).
		SetCustomCell(4, 2, game.Away.Stats.Team.Top, tcell.ColorWhite, tview.AlignCenter).
		SetCustomIntCell(5, 0, game.Home.Timeouts, tcell.ColorWhite, tview.AlignCenter).
		SetCustomIntCell(5, 2, game.Away.Timeouts, tcell.ColorWhite, tview.AlignCenter)
}
