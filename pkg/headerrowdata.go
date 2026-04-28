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

	if header == nil {
		return nil, errors.New("header is nil")
	}
	headers := header.GetCells()

	if len(columndata) != len(headers) {
		return nil, fmt.Errorf("row index %d column data has %d columns but does not have the expected headers %d", rowindex, len(columndata), len(headers))
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

func (hrd *HeaderRowData) GetValStringArray(name string, sep string) ([]string, error) {

	data, err := hrd.GetValData(name)
	if err != nil {
		return []string{}, err
	}
	val, ok := data.(string)
	if !ok {
		return []string{}, errors.New("data was not a string")
	}

	return strings.Split(val, sep), nil
}

func (hrd *HeaderRowData) GetValInt(name string) (int, error) {
	data, err := hrd.GetValString(name)
	if err != nil {
		return -1, err
	}
	return strconv.Atoi(data)
}

func (hrd *HeaderRowData) GetIValData(columnindex int) (any, error) {

	data, ok := hrd.IndexedRawData[columnindex]
	if !ok {
		return nil, fmt.Errorf("could not find data at %d ", columnindex)
	}

	return data, nil
}

func (hrd *HeaderRowData) GetIValString(columnindex int) (string, error) {
	data, err := hrd.GetIValData(columnindex)
	if err != nil {
		return "", err
	}
	val, ok := data.(string)
	if !ok {
		return "", errors.New("data was not a string")
	}

	return val, nil
}

func (hrd *HeaderRowData) GetIValStringArray(columnindex int, sep string) ([]string, error) {

	data, err := hrd.GetIValData(columnindex)
	if err != nil {
		return []string{}, err
	}
	val, ok := data.(string)
	if !ok {
		return []string{}, errors.New("data was not a string")
	}

	return strings.Split(val, sep), nil
}

func (hrd *HeaderRowData) GetIValInt(columnindex int) (int, error) {
	data, err := hrd.GetIValString(columnindex)
	if err != nil {
		return -1, err
	}
	return strconv.Atoi(data)
}
