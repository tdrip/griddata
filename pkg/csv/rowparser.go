package csv

import (
	"errors"
	"io"

	gd "github.com/tdrip/griddata/pkg"
	igrid "github.com/tdrip/griddata/pkg/interfaces"
)

//RowParser parses a csv row by row
type RowParser struct {

	// inherit from the row parser
	igrid.IRowParser
}

//CreateRowParser creates the row parser
func CreateRowParser() *RowParser {
	csvsource := RowParser{}
	return &csvsource
}

//Parse parse the data source
func (rd *RowParser) Parse(parent igrid.IParser, data igrid.IDataSource) error {

	// convert the idatasource to what we expect which is a CSV File
	csvdata := data.(*CSVFile)

	// We need a GD Parser for the logging
	gdp := parent.(*gd.Parser)

	if csvdata != nil {
		row := 0
		pass := 1
		for {
			record, err := csvdata.Reader.Read()
			if err == io.EOF {
				gdp.Logger.LogDebug("Parse", "End of file")
				break
			}
			if err != nil {
				gdp.Logger.LogErrorE("Parse", err)
				break
			} else {
				gdp.Logger.LogDebug("Parse", record)
				rd := CreateRowData(row, pass, record)

				for _, cell := range rd.GetCells() {
					gdp.Logger.LogDebugf("Parse", "Cell %v", cell)
				}

			}
			row++
		}

		return nil

	}

	return errors.New("Data source provided was not of type CSV File or was nil")

}
