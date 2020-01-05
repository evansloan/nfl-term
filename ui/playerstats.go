package ui

import (
	"sort"

	"github.com/evansloan/nfl-term/api"
)

// PlayerStats represents a UI element that displays
// the player stats from a specific NFL game.
// Can display passing/rushing/receiving stats.
//
// Inherits from GenericTable
type PlayerStats struct {
	*GenericTable
	SetStats func(game *api.Game)
}

// NewPlayerStats creates a new PlayerStats UI element
//
// Valid statType values:
//		0 = passing
//		1 = rushing
//		2 = receiving
//
func NewPlayerStats(title string, statType int) *PlayerStats {
	p := &PlayerStats{
		GenericTable: NewGenericTable(title),
	}

	if statType == 0 {
		p.setHandler(p.setPassStats)
	} else if statType == 1 {
		p.setHandler(p.setRushStats)
	} else if statType == 2 {
		p.setHandler(p.setRecvStats)
	}

	return p
}

// setHandler sets the function to be called when PlayerStats.SetStats() is called
func (p *PlayerStats) setHandler(handler func(*api.Game)) {
	p.SetStats = handler
}

// setPassStats populates the UI element with passing stats
// from a specific NFL game.
func (p *PlayerStats) setPassStats(game *api.Game) {
	p.Clear()

	p.SetHeaderCell(0, 0, "Name").
		SetHeaderCell(0, 1, "Comp").
		SetHeaderCell(0, 2, "Att").
		SetHeaderCell(0, 3, "Yards").
		SetHeaderCell(0, 4, "TDs").
		SetHeaderCell(0, 5, "Ints")

	stats := sortPlayers(game.Home.Stats.Passing, game.Away.Stats.Passing)

	for i, player := range stats {
		p.SetTextCell(i+1, 0, player.Name).
			SetIntCell(i+1, 1, player.Comp).
			SetIntCell(i+1, 2, player.Att).
			SetIntCell(i+1, 3, player.Yards).
			SetIntCell(i+1, 4, player.Tds).
			SetIntCell(i+1, 5, player.Ints)
	}
}

// setRushStats populates the UI element with rushing stats
// from a specific NFL game.
func (p *PlayerStats) setRushStats(game *api.Game) {
	p.Clear()

	p.SetHeaderCell(0, 0, "Name").
		SetHeaderCell(0, 1, "Att").
		SetHeaderCell(0, 2, "Yards").
		SetHeaderCell(0, 3, "TDs")

	stats := sortPlayers(game.Home.Stats.Rushing, game.Away.Stats.Rushing)

	for i, player := range stats {
		p.SetTextCell(i+1, 0, player.Name).
			SetIntCell(i+1, 1, player.Att).
			SetIntCell(i+1, 2, player.Yards).
			SetIntCell(i+1, 3, player.Tds)
	}
}

// setRecvStats populates the UI element with receivings stats
// from a specific NFL game.
func (p *PlayerStats) setRecvStats(game *api.Game) {
	p.Clear()

	p.SetHeaderCell(0, 0, "Name").
		SetHeaderCell(0, 1, "Rec").
		SetHeaderCell(0, 2, "Yards").
		SetHeaderCell(0, 3, "TDs")

	stats := sortPlayers(game.Home.Stats.Receiving, game.Away.Stats.Receiving)

	for i, player := range stats {
		p.SetTextCell(i+1, 0, player.Name).
			SetIntCell(i+1, 1, player.Rec).
			SetIntCell(i+1, 2, player.Yards).
			SetIntCell(i+1, 3, player.Tds)
	}

}

// sortPlayers sorts players by passing/rushing/receiving yards
// in descending order.
func sortPlayers(home map[string]api.Player, away map[string]api.Player) []api.Player {
	var stats []api.Player

	for _, player := range home {
		stats = append(stats, player)
	}
	for _, player := range away {
		stats = append(stats, player)
	}

	sort.Slice(stats[:], func(i, j int) bool {
		return stats[i].Yards > stats[j].Yards
	})

	return stats
}
