package data

import (
	"errors"
	"strconv"
	"strings"

	grid "github.com/tdrip/griddata/pkg/grid"
	igrid "github.com/tdrip/griddata/pkg/grid/interfaces"
)

// Row This represents a row of data
type RowData struct {
	grid.Row
}

func NewRowData(rowindex int, pass int) *RowData {
	row := grid.NewRow(rowindex, pass)

	rd := RowData{
		Row: *row,
	}
	return &rd
}

func GetRowData(row igrid.IRow) (*RowData, error) {
	rd, ok := row.(*RowData)
	if !ok {
		return nil, errors.New("data was not Row Data")
	}

	return rd, nil
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

		pnt := grid.NewPoint(rowindex, columnindex)

		// csv is always srting so we parse the cells as such
		cell := grid.NewStringCell(pnt, columndata[columnindex])

		// add the cell
		rd.AddCell(cell)
	}

	return rd
}
