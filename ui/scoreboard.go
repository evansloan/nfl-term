package ui

import (
	"github.com/evansloan/nfl-term/api"
)

// Scoreboard represents a UI element that shows
// the scores of an NFL game quarter by quarter
type Scoreboard struct {
	*DefaultTable
}

// NewScoreboard creates a new Scoreboard element
func NewScoreboard() *Scoreboard {
	return &Scoreboard{
		DefaultTable: NewGenericTable("Scoreboard"),
	}
}

// SetScores populates the scoreboard with the current
// scores of a game.
func (s *Scoreboard) SetScores(game *api.Game) {
	s.Clear()
	s.SetHeaderCell(0, 0, "Team").
		SetHeaderCell(0, 1, "Q1").
		SetHeaderCell(0, 2, "Q2").
		SetHeaderCell(0, 3, "Q3").
		SetHeaderCell(0, 4, "Q4").
		SetHeaderCell(0, 5, "OT").
		SetHeaderCell(0, 6, "F")

	s.SetColorCell(1, 0, game.Home.Abbr, teamColors[game.Home.Abbr]).
		SetIntCell(1, 1, game.Home.Score.Q1).
		SetIntCell(1, 2, game.Home.Score.Q2).
		SetIntCell(1, 3, game.Home.Score.Q3).
		SetIntCell(1, 4, game.Home.Score.Q4).
		SetIntCell(1, 5, game.Home.Score.OT).
		SetIntCell(1, 6, game.Home.Score.Final)

	s.SetColorCell(2, 0, game.Away.Abbr, teamColors[game.Away.Abbr]).
		SetIntCell(2, 1, game.Away.Score.Q1).
		SetIntCell(2, 2, game.Away.Score.Q2).
		SetIntCell(2, 3, game.Away.Score.Q3).
		SetIntCell(2, 4, game.Away.Score.Q4).
		SetIntCell(2, 5, game.Away.Score.OT).
		SetIntCell(2, 6, game.Away.Score.Final)
}
