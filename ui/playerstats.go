package ui

import (
	"sort"

	"github.com/evansloan/nfl-term/api"
)

type StatType int

const (
	Passing   StatType = 0
	Rushing   StatType = 1
	Receiving StatType = 2
	Defense   StatType = 3
	Kicking   StatType = 4
	Punting   StatType = 5
	KickRet   StatType = 6
	PuntRet   StatType = 7
)

// PlayerStats represents a UI element that displays
// the player stats from a specific NFL game.
// Can display passing/rushing/receiving stats.
//
// Inherits from GenericTable
type PlayerStats struct {
	*DefaultTable
	StatType StatType
	SetStats func(game *api.Game)
}

// NewPlayerStats creates a new PlayerStats UI element
func NewPlayerStats(title string, statType StatType) *PlayerStats {
	p := &PlayerStats{
		DefaultTable: NewGenericTable(title),
		StatType:     statType,
	}

	p.SetFixed(1, 1)

	if statType == Passing {
		p.setHandler(p.setPassStats)
	} else if statType == Rushing {
		p.setHandler(p.setRushStats)
	} else if statType == Receiving {
		p.setHandler(p.setRecvStats)
	} else if statType == Defense {
		p.setHandler(p.setDefStats)
	} else if statType == Kicking {
		p.setHandler(p.setKickStats)
	} else if statType == Punting {
		p.setHandler(p.setPuntStats)
	} else if statType == KickRet || statType == PuntRet {
		p.setHandler(p.setReturnStats)
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

	stats := p.sortPlayers(game.Home.Stats.Passing, game.Away.Stats.Passing)

	for i, player := range stats {
		p.SetTextCell(i+1, 0, player.Name).
			SetIntCell(i+1, 1, player.Comp).
			SetIntCell(i+1, 2, player.Att).
			SetIntCell(i+1, 3, player.Yards).
			SetIntCell(i+1, 4, player.Tds).
			SetIntCell(i+1, 5, player.IntsO)
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

	stats := p.sortPlayers(game.Home.Stats.Rushing, game.Away.Stats.Rushing)

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

	stats := p.sortPlayers(game.Home.Stats.Receiving, game.Away.Stats.Receiving)

	for i, player := range stats {
		p.SetTextCell(i+1, 0, player.Name).
			SetIntCell(i+1, 1, player.Rec).
			SetIntCell(i+1, 2, player.Yards).
			SetIntCell(i+1, 3, player.Tds)
	}
}

// setDefStats populates the UI element with receivings stats
// from a specific NFL game.
func (p *PlayerStats) setDefStats(game *api.Game) {
	p.Clear()

	p.SetHeaderCell(0, 0, "Name").
		SetHeaderCell(0, 1, "Tkls").
		SetHeaderCell(0, 2, "Asts").
		SetHeaderCell(0, 3, "Sacks").
		SetHeaderCell(0, 4, "Ints").
		SetHeaderCell(0, 5, "FFum")

	stats := p.sortPlayers(game.Home.Stats.Defense, game.Away.Stats.Defense)

	for i, player := range stats {
		p.SetTextCell(i+1, 0, player.Name).
			SetIntCell(i+1, 1, player.Tkl).
			SetIntCell(i+1, 2, player.Ast).
			SetIntCell(i+1, 3, player.Sacks).
			SetIntCell(i+1, 4, player.IntsD).
			SetIntCell(i+1, 5, player.Ffum)
	}
	p.ScrollToBeginning()
}

func (p *PlayerStats) setKickStats(game *api.Game) {
	p.Clear()

	p.SetHeaderCell(0, 0, "Name").
		SetHeaderCell(0, 1, "FG Made").
		SetHeaderCell(0, 2, "FG Att").
		SetHeaderCell(0, 3, "XP Made").
		SetHeaderCell(0, 4, "XP Att")

	stats := p.sortPlayers(game.Home.Stats.Kicking, game.Away.Stats.Kicking)

	for i, player := range stats {
		p.SetTextCell(i+1, 0, player.Name).
			SetIntCell(i+1, 1, player.FGMade).
			SetIntCell(i+1, 2, player.FGAtt).
			SetIntCell(i+1, 3, player.XPMade).
			SetIntCell(i+1, 4, player.XPAtt)
	}
	p.ScrollToBeginning()
}

func (p *PlayerStats) setPuntStats(game *api.Game) {
	p.Clear()

	p.SetHeaderCell(0, 0, "Name").
		SetHeaderCell(0, 1, "Punts").
		SetHeaderCell(0, 2, "Yards").
		SetHeaderCell(0, 3, "Inside 20")

	stats := p.sortPlayers(game.Home.Stats.Punting, game.Away.Stats.Punting)

	for i, player := range stats {
		p.SetTextCell(i+1, 0, player.Name).
			SetIntCell(i+1, 1, player.Punts).
			SetIntCell(i+1, 2, player.Yards).
			SetIntCell(i+1, 3, player.Inside20)
	}
	p.ScrollToBeginning()
}

func (p *PlayerStats) setReturnStats(game *api.Game) {
	p.Clear()

	p.SetHeaderCell(0, 0, "Name").
		SetHeaderCell(0, 1, "Returns").
		SetHeaderCell(0, 2, "Avg").
		SetHeaderCell(0, 3, "Yards").
		SetHeaderCell(0, 4, "TDs").
		SetHeaderCell(0, 5, "Long")

	var stats []api.Player

	if p.StatType == KickRet {
		stats = p.sortPlayers(game.Home.Stats.KickRet, game.Away.Stats.KickRet)
	} else if p.StatType == PuntRet {
		stats = p.sortPlayers(game.Home.Stats.PuntRet, game.Away.Stats.PuntRet)
	}

	for i, player := range stats {
		p.SetTextCell(i+1, 0, player.Name).
			SetIntCell(i+1, 1, player.Returns).
			SetIntCell(i+1, 2, player.Avg).
			SetIntCell(i+1, 3, player.Yards).
			SetIntCell(i+1, 4, player.Tds).
			SetIntCell(i+1, 5, player.Long)
	}
	p.ScrollToBeginning()
}

// sortPlayers sorts players by passing/rushing/receiving yards
// in descending order.
func (p *PlayerStats) sortPlayers(home, away map[string]api.Player) []api.Player {
	var stats []api.Player

	for _, player := range home {
		stats = append(stats, player)
	}
	for _, player := range away {
		stats = append(stats, player)
	}

	sort.Slice(stats[:], func(i, j int) bool {
		if p.StatType == Passing || p.StatType == Rushing || p.StatType == Receiving {
			return stats[i].Yards > stats[j].Yards
		} else if p.StatType == Defense {
			return stats[i].Tkl > stats[j].Tkl
		} else if p.StatType == Kicking {
			return stats[i].FGMade > stats[j].FGMade
		} else if p.StatType == Punting {
			return stats[i].Punts > stats[j].Punts
		} else if p.StatType == KickRet || p.StatType == PuntRet {
			return stats[i].Att > stats[j].Att
		}
		return []rune(stats[i].Name)[0] < []rune(stats[j].Name)[0]
	})

	return stats
}
