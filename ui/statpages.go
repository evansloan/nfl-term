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

func newStatPage(stats []*PlayerStats, cols ...int) *statPage {
	s := &statPage{
		Grid:     tview.NewGrid(),
		Stats:    stats,
		selected: -1,
	}

	s.SetColumns(cols...)

	for i, item := range stats {
		item.SetSelectable(true, false)
		s.AddItem(item, 0, i, 3, 1, 0, 0, true)
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
	stPage     *statPage
	ActivePage *statPage
}

func NewStatPages() *StatPages {
	s := &StatPages{
		Pages: tview.NewPages(),
	}

	oStats := []*PlayerStats{
		NewPlayerStats("Passing", Passing),
		NewPlayerStats("Rushing", Rushing),
		NewPlayerStats("Receiving", Receiving),
	}
	dStats := []*PlayerStats{
		NewPlayerStats("Defense", Defense),
	}
	stStats := []*PlayerStats{
		NewPlayerStats("Kicking", Kicking),
		NewPlayerStats("Punting", Punting),
		NewPlayerStats("Kick return", KickRet),
		NewPlayerStats("Punt return", PuntRet),
	}
	s.oPage = newStatPage(oStats, 40, 30, 30, 0)
	s.dPage = newStatPage(dStats, 60, 0)
	s.stPage = newStatPage(stStats, 30, 30, 30, 30, 0)

	s.AddPage("opage", s.oPage, true, true)
	s.AddPage("dpage", s.dPage, true, false)
	s.AddPage("stpage", s.stPage, true, false)
	s.ActivePage = s.oPage

	return s
}

func (s *StatPages) SetActive(pageName string) {
	if pageName == "opage" {
		s.ActivePage = s.oPage
	} else if pageName == "dpage" {
		s.ActivePage = s.dPage
	} else if pageName == "stpage" {
		s.ActivePage = s.stPage
	}
}

func (s *StatPages) SetStats(game *api.Game) {
	for _, stat := range s.oPage.Stats {
		stat.SetStats(game)
	}
	for _, stat := range s.dPage.Stats {
		stat.SetStats(game)
	}
	for _, stat := range s.stPage.Stats {
		stat.SetStats(game)
	}
}
