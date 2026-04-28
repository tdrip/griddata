package grid

import (
	"errors"
	"strconv"
	"strings"

	igrid "github.com/tdrip/griddata/pkg/interfaces"
)

// RowData This represents a row of data
type RowData struct {
	igrid.IRow

	// Index of the row
	Index igrid.IIndex

	// Number of passes over the row
	Pass int

	// Parsed Cell Data
	// index these!
	Cells []igrid.ICell
}

// NewRowData creates a default row data struct
func NewRowData(rowindex int, pass int) *RowData {
	rd := RowData{Pass: pass}

	// x,y point doesn;t matter

	// set the index
	rd.SetIndex(JustXIndex(rowindex))

	return &rd
}

// GetIndex Gets the index for the row
func (rd *RowData) GetIndex() igrid.IIndex {
	return rd.Index
}

// SetIndex Sets the index for the row
func (rd *RowData) SetIndex(index igrid.IIndex) {
	rd.Index = index
}

// Matches Matches the index passed in against the index for the row
func (rd *RowData) Matches(index igrid.IIndex) bool {
	return rd.GetIndex().GetLocation().Match(index.GetLocation())
}

// GetCells Gets the cells for the row
func (rd *RowData) GetCells() []igrid.ICell {
	return rd.Cells
}

// SetCells Sets the cells for the row
func (rd *RowData) SetCells(cells []igrid.ICell) {
	rd.Cells = cells
}

// AddCell Add a cells to the row
func (rd *RowData) AddCell(cell igrid.ICell) {
	cells := rd.Cells
	cells = append(cells, cell)
	rd.Cells = cells
}

func (rd *RowData) GetValData(columnindex int) (any, error) {
	if columnindex < 0 || columnindex >= len(rd.Cells) {
		return nil, errors.New("columns out of range")
	}
	return rd.Cells[columnindex].GetData(), nil
}

func (rd *RowData) GetValString(columnindex int) (string, error) {
	data, err := rd.GetValData(columnindex)
	if err != nil {
		return "", err
	}
	val, ok := data.(string)
	if !ok {
		return "", errors.New("data was not a string")
	}

	return val, nil
}

func (rd *RowData) GetValStringArray(columnindex int, sep string) ([]string, error) {

	data, err := rd.GetValData(columnindex)
	if err != nil {
		return []string{}, err
	}
	val, ok := data.(string)
	if !ok {
		return []string{}, errors.New("data was not a string")
	}

	return strings.Split(val, sep), nil
}

func (rd *RowData) GetValInt(columnindex int) (int, error) {
	data, err := rd.GetValString(columnindex)
	if err != nil {
		return -1, err
	}
	return strconv.Atoi(data)
}

// FillRowStringData creates a row data from a string data array
func FillRowStringData(rowindex int, pass int, columndata []string) *RowData {

	// number of passes and the row index
	rd := NewRowData(rowindex, pass)

	for columnindex := 0; columnindex < len(columndata); columnindex++ {

		pnt := NewPoint(rowindex, columnindex)
		// csv is always srting so we parse the cells as such
		cell := NewStringCell(pnt, columndata[columnindex])

		// add the cell
		rd.AddCell(cell)
	}

	return rd
}
