package csv

import (
	"errors"
	"fmt"
	"io"

	gd "github.com/tdrip/griddata/pkg"
	igrid "github.com/tdrip/griddata/pkg/interfaces"
)

//RowParsingOptions number of passes etc
type RowParsingOptions struct {

	// inherit from the row IRowParsingOptions
	igrid.IRowParsingOptions

	TotalPasses int `json:"totalpasses"`

	HeaderRowIndex int `json:"headerrowindex"`
}

//RowParser parses a csv row by row
type RowParser struct {

	// inherit from the row parser
	igrid.IRowParser

	Options *RowParsingOptions

	Actions map[string]igrid.IRowAction
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
func (parser *RowParser) Parse(parent igrid.IParser, data igrid.IDataSource) error {

	// convert the idatasource to what we expect which is a CSV File
	csvdata := data.(*CSVFile)

	// We need a GD Parser for the logging
	gdp := parent.(*gd.Parser)

	// We need a GD Parser for the logging
	opts := parser.GetOptions()
	options := opts.(*RowParsingOptions)

	if csvdata != nil {
		row := 0
		pass := options.TotalPasses
		for {
			record, err := csvdata.Reader.Read()
			if err == io.EOF {
				gdp.Logger.Debugf("%s - %s", "Parse", "End of file")
				break
			}
			if err != nil {
				gdp.Logger.Error("Parse", err)
				return err
			} else {
				gdp.Logger.Debugf("%s - %v", "Parse", record)

				if row == options.HeaderRowIndex {

				} else {

					//create row data
					rd := CreateRowData(row, pass, record)

					// Get cells from the row
					for _, cell := range rd.GetCells() {
						// print the cells that we read
						gdp.Logger.Debugf("%s - %v", "Parse", cell)
					}

					// get the actions
					for _, action := range parser.GetActions() {

						// perform action on teh row data
						action.PeformAction(rd)
					}
				}

			}
			row++
		}

		return nil

	}

	return errors.New("data source provided was not of type CSV File or was nil")

}

//GetOptions Get the options for the row parser
func (rd *RowParser) GetOptions() igrid.IRowParsingOptions {
	return rd.Options
}

//SetOptions Set the options for the row parser
func (rd *RowParser) SetOptions(options igrid.IRowParsingOptions) {
	rd.Options = options.(*RowParsingOptions)
}

// actions for the row
func (rd *RowParser) GetActions() []igrid.IRowAction {
	return nil
}

func (rd *RowParser) SetActions(data map[string]igrid.IRowAction) {
	rd.Actions = data
}

func (rd *RowParser) AddAction(action igrid.IRowAction) {
	data := rd.Actions
	data[action.GetId()] = action
	rd.Actions = data
}

////////////////////////////////////
// ROW OPTIONS
////////////////////////////////////

//Defaults
func (rpo *RowParsingOptions) Defaults() {
	//Only pass over the row once
	rpo.TotalPasses = 1

	// default to no header
	rpo.HeaderRowIndex = -1
}

//String the reable version of the options
func (rpo *RowParsingOptions) String() string {
	return fmt.Sprintf("Total Row Passes: %d", rpo.TotalPasses)
}
