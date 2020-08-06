package csv

import (
	"errors"
	"io"

	gd "github.com/tdrip/griddata/pkg"
	igrid "github.com/tdrip/griddata/pkg/interfaces"
)

//RowParsingOptions number of passes etc
type RowParsingOptions struct {

	// inherit from the row IRowParsingOptions
	igrid.IRowParsingOptions

	TotalPasses int
}

//RowParser parses a csv row by row
type RowParser struct {

	// inherit from the row parser
	igrid.IRowParser

	Options *RowParsingOptions
}

//CreateRowParser creates the row parser
func CreateRowParser() *RowParser {
	csvsource := &RowParser{}
	//Options
	opts := &RowParsingOptions{}
	opts.Defaults()
	csvsource.SetOptions(opts)
	return csvsource
}

//Parse parse the data source
func (rd *RowParser) Parse(parent igrid.IParser, data igrid.IDataSource) error {

	// convert the idatasource to what we expect which is a CSV File
	csvdata := data.(*CSVFile)

	// We need a GD Parser for the logging
	gdp := parent.(*gd.Parser)

	// We need a GD Parser for the logging
	opts := rd.GetOptions()
	options := opts.(*RowParsingOptions)

	if csvdata != nil {
		row := 0
		pass := options.TotalPasses
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
					// print the cells that we read
					gdp.Logger.LogDebugf("Parse", "%s", cell)
				}

			}
			row++
		}

		return nil

	}

	return errors.New("Data source provided was not of type CSV File or was nil")

}

//GetOptions Get the options for the row parser
func (rd *RowParser) GetOptions() igrid.IRowParsingOptions {
	return rd.Options
}

//SetOptions Set the options for the row parser
func (rd *RowParser) SetOptions(options igrid.IRowParsingOptions) {
	rd.Options = options.(*RowParsingOptions)
}

//Defaults
func (rpo *RowParsingOptions) Defaults() {
	//Only pass over the row once
	rpo.TotalPasses = 1
}

//String the reable version of the options
func (rpo *RowParsingOptions) String() string {
	return ""
}
