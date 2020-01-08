package ui

import (
	"strconv"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

// GenericTable is the default tview.Table configuration
// for the application.
//
// Inherits from tview.Table
type GenericTable struct {
	*tview.Table
}

// NewGenericTable creates a new Table element and
// applies basic styling.
func NewGenericTable(title string) *GenericTable {
	g := &GenericTable{
		Table: tview.NewTable(),
	}
	g.SetTitle(title).
		SetBorder(true)
	g.SetBorders(false)
	return g
}

// SetTextCell creates a basic cell that contains text within
// a table.
//
// Is the same as tview.Table.SetCellSimple() but returns *GenericTable
// instead of *tview.Table.
func (g *GenericTable) SetTextCell(row int, col int, text string) *GenericTable {
	g.SetCellSimple(row, col, text)
	return g
}

// SetHeaderCell creates a cell that has a darker color than a standard
// table cell.
// Is used to keep table headers/categories uniform.
func (g *GenericTable) SetHeaderCell(row int, col int, text string) *GenericTable {
	g.SetCell(row, col, &tview.TableCell{Text: text, Color: tcell.ColorLightGray})
	return g
}

// SetColorCell creates a cell that contains text of a specific color
// within a table.
func (g *GenericTable) SetColorCell(row int, col int, text string, color tcell.Color) *GenericTable {
	g.SetCell(row, col, &tview.TableCell{Text: text, Color: color})
	return g
}

// SetIntCell converts integers into text that can be placed within
// a cell within a table.
func (g *GenericTable) SetIntCell(row int, col int, text int) *GenericTable {
	g.SetTextCell(row, col, strconv.Itoa(text))
	return g
}

// SetTextCell creates a customizable cell that contains text within
// a table.
//
// Is the same as tview.Table.SetCell() but returns *GenericTable
// instead of *tview.Table.
func (g *GenericTable) SetCustomCell(row, col int, text string, color tcell.Color, align int) *GenericTable {
	cell := &tview.TableCell{Text: text, Color: color, Align: align}
	g.SetCell(row, col, cell)
	return g
}
