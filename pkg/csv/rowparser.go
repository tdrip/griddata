package csv

import (
	"errors"
	"io"
	"log"

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

	csvdata := data.(*CSVFile)
	gdp := parent.(*gd.GDParser)

	if csvdata != nil {
		for {

			record, err := csvdata.Reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			gdp.Logger.LogInfo("Parse", record)
		}

		return nil

	}

	return errors.New("Data source was not of type CSV Files or was nil")

}
