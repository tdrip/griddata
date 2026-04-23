package grid

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// RowData This represents a row of data
type HeaderRowData struct {
	RowData

	IndexedRawData map[int]any
	NameLookUp     map[string]int
	Header         *RowData
}

// CreateRowData Creates a default row data struct
func CreateHeaderRowData(row int, pass int, header *RowData) *HeaderRowData {
	rd := HeaderRowData{RowData: *CreateRowData(row, pass)}
	rd.Header = header
	return &rd
}

// FillHeaderRowStringData creates a row data from a string data array
func FillHeaderRowStringData(rowindex int, pass int, columndata []string, header *RowData) (*HeaderRowData, error) {

	headers := header.GetCells()
	if len(columndata) != len(headers) {
		return nil, errors.New("headers does not have the same number of columns as header")
	}
	// number of passes and the row index
	rd := CreateHeaderRowData(rowindex, pass, header)
	indexeddata := make(map[int]any, len(columndata))
	namelookup := make(map[string]int, len(columndata))

	for columnindex := 0; columnindex < len(columndata); columnindex++ {

		pnt := CreatePoint(rowindex, columnindex)
		// csv is always srting so we parse the cells as such
		cell := CreateStringCell(pnt, columndata[columnindex])

		head := headers[columnindex].GetData()
		headv, ok := head.(string)
		if !ok {
			return nil, errors.New("headers does not have the same number of columns as header")
		}
		namelookup[headv] = columnindex
		indexeddata[namelookup[headv]] = columndata[columnindex]

		// add the cell
		rd.AddCell(cell)
	}

	rd.IndexedRawData = indexeddata
	rd.NameLookUp = namelookup

	return rd, nil
}

func (hrd *HeaderRowData) GetValData(name string) (any, error) {
	columnindex, ok := hrd.NameLookUp[name]
	if !ok {
		columnindex, ok = hrd.NameLookUp[strings.ToLower(name)]
		if !ok {
			found := false
			for lname, index := range hrd.NameLookUp {
				if strings.EqualFold(lname, name) {
					columnindex = index
					found = true
					break
				}
			}
			if !found {
				return nil, fmt.Errorf("could not find header column with name %s", name)
			}
		}
	}

	data, ok := hrd.IndexedRawData[columnindex]
	if !ok {
		return nil, fmt.Errorf("could not find data at %d for %s", columnindex, name)
	}

	return data, nil
}

func (hrd *HeaderRowData) GetValString(name string) (string, error) {

	data, err := hrd.GetValData(name)
	if err != nil {
		return "", err
	}
	val, ok := data.(string)
	if !ok {
		return "", errors.New("data was not a string")
	}

	return val, nil
}

func (hrd *HeaderRowData) GetValInt(name string) (int, error) {
	data, err := hrd.GetValString(name)
	if err != nil {
		return -1, err
	}
	return strconv.Atoi(data)
}
