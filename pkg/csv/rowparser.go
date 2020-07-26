package csv

import (
	"errors"
	"io"

	gd "github.com/tdrip/griddata/pkg"
	igrid "github.com/tdrip/griddata/pkg/interfaces"
)

//CSVRowParser parses a csv row by row
type CSVRowParser struct {

	// inherit from the row parser
	igrid.IRowParser
}

func CreateCSVRowParser() *CSVRowParser {
	csvsource := CSVRowParser{}
	return &csvsource
}

//Parse parse the data source
func (rd *CSVRowParser) Parse(parent igrid.IParser, data igrid.IDataSource) error {

	// convert the idatasource to what we expect which is a CSV File
	csvdata := data.(*CSVFile)

	// We need a GD Parser for the logging
	gdp := parent.(*gd.GDParser)

	if csvdata != nil {
		row := 0
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
				gdp.Logger.LogInfo("Parse", record)
			}
			row++
		}

		return nil

	}

	return errors.New("Data source provided was not of type CSV File or was nil")

}
