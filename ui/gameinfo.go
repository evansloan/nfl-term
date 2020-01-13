package ui

import (
	"strconv"
	"strings"

	"github.com/evansloan/nfl-term/api"
)

// GameInfo represents a UI element that displays
// relevant information of an NFL game.
//
// Inherits from DefaultTable
type GameInfo struct {
	*DefaultTable
}

// NewGameInfo creates a new GameInfo UI element
func NewGameInfo() *GameInfo {
	return &GameInfo{
		DefaultTable: NewDefaultTable("Game Info"),
	}
}

// SetInfo applies game information to a GameInfo UI element
func (g *GameInfo) SetInfo(game *api.Game) {
	g.Clear()

	if strings.Contains(game.Qtr, "final") || game.Qtr == "Final" {
		game.Clock = ""
		game.Pos = ""
	}

	down := strconv.Itoa(game.Down)
	togo := strconv.Itoa(game.ToGo)
	downStr := ""
	if game.Down == 1 {
		downStr = down + "st and " + togo
	} else if game.Down == 2 {
		downStr = down + "nd and " + togo
	} else if game.Down == 3 {
		downStr = down + "rd and " + togo
	} else if game.Down == 4 {
		downStr = down + "th and " + togo
	}

	g.SetHeaderCell(0, 0, "Quarter").
		SetHeaderCell(1, 0, "Time").
		SetHeaderCell(2, 0, "Down").
		SetHeaderCell(3, 0, "Possession").
		SetHeaderCell(4, 0, "Yard line")

	g.SetTextCell(0, 1, strings.Title(game.Qtr)).
		SetTextCell(1, 1, game.Clock).
		SetTextCell(2, 1, downStr).
		SetTextCell(3, 1, game.Pos).
		SetTextCell(4, 1, game.Yl)
}
