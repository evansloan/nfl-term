package ui

import (
	"github.com/evansloan/nfl-term/api"
	"github.com/rivo/tview"
)

type statPage struct {
	*tview.Grid
	Stats    []*PlayerStats
	selected int
}

func newStatPage(stats []*PlayerStats, colSpan int) *statPage {
	s := &statPage{
		Grid:     tview.NewGrid(),
		Stats:    stats,
		selected: -1,
	}

	s.SetColumns(40, 30, 30, 0)

	for i, item := range stats {
		item.SetSelectable(true, false)
		s.AddItem(item, 0, i, 3, colSpan, 0, 0, true)
	}

	return s
}

func (s *statPage) SelectNext() *PlayerStats {
	if s.selected+1 == len(s.Stats) {
		s.selected = 0
	} else {
		s.selected = s.selected + 1
	}
	return s.Stats[s.selected]
}

type StatPages struct {
	*tview.Pages
	oPage      *statPage
	dPage      *statPage
	ActivePage *statPage
}

func NewStatPages() *StatPages {
	s := &StatPages{
		Pages: tview.NewPages(),
	}

	oStats := []*PlayerStats{
		NewPlayerStats("Passing", 0),
		NewPlayerStats("Rushing", 1),
		NewPlayerStats("Receiving", 2),
	}
	dStats := []*PlayerStats{
		NewPlayerStats("Defense", 3),
	}
	s.oPage = newStatPage(oStats, 1)
	s.dPage = newStatPage(dStats, 2)

	s.AddPage("opage", s.oPage, true, true)
	s.AddPage("dpage", s.dPage, true, false)
	s.ActivePage = s.oPage

	return s
}

func (s *StatPages) SetActive(pageName string) {
	if pageName == "opage" {
		s.ActivePage = s.oPage
	} else {
		s.ActivePage = s.dPage
	}
}

func (s *StatPages) SetStats(game *api.Game) {
	for _, stat := range s.oPage.Stats {
		stat.SetStats(game)
	}
	for _, stat := range s.dPage.Stats {
		stat.SetStats(game)
	}
}
