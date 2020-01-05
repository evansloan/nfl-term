package ui

import (
	"github.com/evansloan/nfl-term/api"
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
	
	t.SetColorCell(0, 1, game.Home.Abbr, teamColors[game.Home.Abbr]).
		SetColorCell(0, 2, game.Away.Abbr, teamColors[game.Away.Abbr])

	t.SetHeaderCell(1, 0, "Pass:").
		SetHeaderCell(2, 0, "Rush:").
		SetHeaderCell(3, 0, "Total:").
		SetHeaderCell(4, 0, "Time of Pos:").
		SetHeaderCell(5, 0, "Timeouts:")

	t.SetIntCell(1, 1, game.Home.Stats.Team.PassYards).
		SetIntCell(1, 2, game.Away.Stats.Team.PassYards).
		SetIntCell(2, 1, game.Home.Stats.Team.RushYards).
		SetIntCell(2, 2, game.Away.Stats.Team.RushYards).
		SetIntCell(3, 1, game.Home.Stats.Team.TotalYards).
		SetIntCell(3, 2, game.Away.Stats.Team.TotalYards).
		SetTextCell(4, 1, game.Home.Stats.Team.Top).
		SetTextCell(4, 2, game.Away.Stats.Team.Top).
		SetIntCell(5, 1, game.Home.Timeouts).
		SetIntCell(5, 2, game.Away.Timeouts)
}