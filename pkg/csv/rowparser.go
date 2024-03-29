package csv

import (
	"errors"
	"fmt"
	"io"

	gd "github.com/tdrip/griddata/pkg"
	igrid "github.com/tdrip/griddata/pkg/interfaces"
)

//RowParsingOptions number of passes etc
type RowProcessorOptions struct {

	// inherit from the row IRowProcessingOptions
	igrid.IDataProcessorOptions

	TotalPasses int `json:"totalpasses"`

	HeaderRowIndex int `json:"headerrowindex"`
}

//RowProcessor parses a csv row by row
type RowProcessor struct {

	// inherit from the row parser
	igrid.IDataProcessor

	Options *RowProcessorOptions

	Actions map[string]igrid.IDataAction
}

//CreateRowProcessor creates the row parser
func CreateRowProcessor() *RowProcessor {
	csvsource := &RowProcessor{}
	//Options
	opts := &RowProcessorOptions{}
	opts.Defaults()
	csvsource.SetOptions(opts)
	empty := make(map[string]igrid.IDataAction)
	csvsource.Actions = empty
	return csvsource
}

//Parse parse the data source
func (rowparser *RowProcessor) Parse(parent igrid.IParser, data igrid.IDataSource) error {

	// convert the idatasource to what we expect which is a CSV File
	csvdata := data.(*CSVFile)

	// We need a GD Parser for the logging
	gdp := parent.(*gd.Parser)

	// We need a GD Parser for the logging
	opts := rowparser.GetOptions()
	options := opts.(*RowProcessorOptions)

	if csvdata != nil {
		row := 0
		pass := options.TotalPasses
		for {
			record, err := csvdata.Reader.Read()
			if err == io.EOF {
				gdp.Logger.Debug("csv parse - End of file")
				break
			}
			if err != nil {
				gdp.Logger.Errorf("csv parse - %v", err)
				return err
			} else {
				gdp.Logger.Debugf("csv parse - Record: %v", record)

				if row == options.HeaderRowIndex {
					gdp.Logger.Debug("csv parse - Header row index")
				} else {

					//create row data
					rd := CreateRowData(row, pass, record)

					// Get cells from the row
					for _, cell := range rd.GetCells() {
						// print the cells that we read
						gdp.Logger.Debugf("csv parse - %v", cell)
					}

					// get the actions
					for _, action := range rowparser.GetActions() {

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
func (rowparser *RowProcessor) GetOptions() igrid.IDataProcessorOptions {
	return rowparser.Options
}

//SetOptions Set the options for the row parser
func (rowparser *RowProcessor) SetOptions(options igrid.IDataProcessorOptions) {
	rowparser.Options = options.(*RowProcessorOptions)
}

// actions for the row
func (rowparser *RowProcessor) GetActions() map[string]igrid.IDataAction {
	return rowparser.Actions
}

func (rowparser *RowProcessor) SetActions(actions []igrid.IDataAction) {
	for _, action := range actions {
		rowparser.AddAction(action)
	}
}

func (rowparser *RowProcessor) AddAction(action igrid.IDataAction) {
	data := rowparser.Actions
	data[action.GetId()] = action
	rowparser.Actions = data
}

func (rowparser *RowProcessor) RemoveAction(id string) {
	data := rowparser.Actions
	delete(data, id)
	rowparser.Actions = data
}

func (rowparser *RowProcessor) ClearActions() {
	empty := make(map[string]igrid.IDataAction)
	rowparser.Actions = empty
}

////////////////////////////////////
// ROW OPTIONS
////////////////////////////////////

//Defaults
func (rpo *RowProcessorOptions) Defaults() {
	//Only pass over the row once
	rpo.TotalPasses = 1

	// default to no header
	rpo.HeaderRowIndex = -1
}

//String the readable version of the options
func (rpo *RowProcessorOptions) String() string {
	return fmt.Sprintf("Total Row Passes: %d", rpo.TotalPasses)
}

//CreateRowData creates a row data from a parsed CSV
func CreateRowData(rowindex int, pass int, columndata []string) *gd.RowData {

	// number of passes and the row index
	rd := gd.CreateRowData(rowindex, pass)

	for columnindex := 0; columnindex < len(columndata); columnindex++ {

		pnt := gd.CreatePoint(rowindex, columnindex)
		// csv is always srting so we parse the cells as such
		cell := gd.CreateStringCell(pnt, columndata[columnindex])

		// add the cell
		rd.AddCell(cell)
	}

	return rd
}
