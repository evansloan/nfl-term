package ui

import (
	"github.com/evansloan/nfl-term/api"
	"github.com/gdamore/tcell"
)

// TeamStats represents a UI element that displays
// basic team stats for an NFL game.
type TeamStats struct {
	*GenericTable
}

// NewTeamStats creates a new TeamStats element.
func NewTeamStats() *TeamStats {
	return &TeamStats{
		GenericTable: NewGenericTable("Team Stats"),
	}
}

// SetStats populates a TeamStats element with stats
// from a specific NFL game.
func (t *TeamStats) SetStats(game *api.Game) {
	t.Clear()

	t.SetColorCell(0, 0, game.Home.Abbr, teamColors[game.Home.Abbr]).
		SetColorCell(0, 2, game.Away.Abbr, teamColors[game.Away.Abbr])

	headerColor := tcell.ColorLightGray
	t.SetCustomCell(1, 1, "Pass", headerColor, 1).
		SetCustomCell(2, 1, "Rush", headerColor, 1).
		SetCustomCell(3, 1, "Total", headerColor, 1).
		SetCustomCell(4, 1, "Time of Pos", headerColor, 1).
		SetCustomCell(5, 1, "Timeouts", headerColor, 1)

	t.SetIntCell(1, 0, game.Home.Stats.Team.PassYards).
		SetIntCell(1, 2, game.Away.Stats.Team.PassYards).
		SetIntCell(2, 0, game.Home.Stats.Team.RushYards).
		SetIntCell(2, 2, game.Away.Stats.Team.RushYards).
		SetIntCell(3, 0, game.Home.Stats.Team.TotalYards).
		SetIntCell(3, 2, game.Away.Stats.Team.TotalYards).
		SetTextCell(4, 0, game.Home.Stats.Team.Top).
		SetTextCell(4, 2, game.Away.Stats.Team.Top).
		SetIntCell(5, 0, game.Home.Timeouts).
		SetIntCell(5, 2, game.Away.Timeouts)
}
