package pt

import (
	"fmt"
	"strings"
)

type Table struct {
	headers []string
	rows    [][]string
}

func NewTable(headers []string) *Table {
	return &Table{
		headers: headers,
		rows:    make([][]string, 0),
	}
}

func (t *Table) AddRow(row []string) {
	t.rows = append(t.rows, row)
}

func (t *Table) Print() {
	tableWidth := len(t.headers)

	// calculate the max. width for each col.
	columnWidths := make([]int, tableWidth)
	for i := 0; i < tableWidth; i++ {
		columnWidths[i] = len(t.headers[i])
	}

	for _, row := range t.rows {
		for i, cell := range row {
			if len(cell) > columnWidths[i] {
				columnWidths[i] = len(cell)
			}
		}
	}

	// print the headers.
	fmt.Println(strings.Join(t.headers, " | "))

	// print the seperator.
	separator := make([]string, tableWidth)
	for i, width := range columnWidths {
		separator[i] = strings.Repeat("-", width)
	}
	fmt.Println(strings.Join(separator, "-+-"))

	// print the rows.
	for _, row := range t.rows {
		cells := make([]string, tableWidth)
		for i, cell := range row {
			cells[i] = cell + strings.Repeat(" ", columnWidths[i]-len(cell))
		}
		fmt.Println(strings.Join(cells, " | "))
	}
}
