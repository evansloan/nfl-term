package ui

import (
	"sort"
	"strconv"

	"github.com/evansloan/nfl-term/api"
)

// ScoreSummary represents a UI element that displays
// a brief description of all scoring plays in a game.
type ScoreSummary struct {
	*GenericTable
}

// NewScoreSummary creates a new ScoreSummary element
func NewScoreSummary() *ScoreSummary {
	return &ScoreSummary{
		GenericTable: NewGenericTable("Score Summary"),
	}
}

// SetScoreSummary populates a ScoreSummary element with
// details of scoring plays.
func (s *ScoreSummary) SetScoreSummary(game *api.Game) {
	s.Clear()

	s.SetHeaderCell(0, 0, "Type").
		SetHeaderCell(0, 1, "Team").
		SetHeaderCell(0, 2, "Quarter").
		SetHeaderCell(0, 3, "Desc")

	// Convert scoring play IDs to ints in order to sort by time scored
	var ids []int
	for id := range game.ScoreSummary {
		intID, _ := strconv.Atoi(id)
		ids = append(ids, intID)
	}
	sort.Ints(ids)

	for i, id := range ids {
		strID := strconv.Itoa(id)
		s.SetTextCell(i+1, 0, game.ScoreSummary[strID].Type).
			SetTextCell(i+1, 1, game.ScoreSummary[strID].Team).
			SetIntCell(i+1, 2, game.ScoreSummary[strID].Qtr).
			SetTextCell(i+1, 3, game.ScoreSummary[strID].Desc)
	}
}
