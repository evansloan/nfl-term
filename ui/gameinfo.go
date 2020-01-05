package ui

import (
	"strconv"
	"strings"

	"github.com/evansloan/nfl-term/api"
)

// GameInfo represents a UI element that displays
// relevant information of an NFL game.
// Inherits from GenericTable
type GameInfo struct {
	*GenericTable
}

// NewGameInfo creates a new GameInfo UI element
func NewGameInfo() *GameInfo {
	return &GameInfo{
		GenericTable: NewGenericTable("Game Info"),
	}
}

// SetInfo applies game information to a GameInfo UI element
func (g *GameInfo) SetInfo(game *api.Game) {
	g.Clear()

	ord := "th"
	if game.Down == 1 {
		ord = "st"
	} else if game.Down == 2 {
		ord = "nd"
	} else if game.Down == 3 {
		ord = "rd"
	}

	g.SetHeaderCell(0, 0, "Quarter:").
		SetHeaderCell(1, 0, "Time:").
		SetHeaderCell(2, 0, "Down:").
		SetHeaderCell(3, 0, "Possession:").
		SetHeaderCell(4, 0, "Yard line:")

	g.SetTextCell(0, 1, strings.Title(game.Qtr)).
		SetTextCell(1, 1, game.Clock).
		SetTextCell(2, 1, strconv.Itoa(game.Down)+ord+" and "+strconv.Itoa(game.ToGo)).
		SetTextCell(3, 1, game.Pos).
		SetTextCell(4, 1, game.Yl)
}
