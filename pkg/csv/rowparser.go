package csv

import (
	"errors"
	"io"

	gd "github.com/tdrip/griddata/pkg"
	igrid "github.com/tdrip/griddata/pkg/interfaces"
)

// Parse parse the data source
func CSVParse(rowparser *gd.RowProcessor, parent igrid.IParser, data igrid.IDataSource) error {

	// convert the idatasource to what we expect which is a CSV File
	csvdata, ok := data.(*CSVFile)
	if !ok {
		return errors.New("data type was not a CSV File")
	}
	// We need a GD Parser for the logging
	opts := rowparser.GetOptions()
	options, ok := opts.(*gd.RowProcessorOptions)
	if !ok {
		return errors.New("options type was not a Row Processor Options")
	}
	var hrd *gd.HeaderRowData
	hrd = nil
	if csvdata != nil {
		row := 0
		pass := options.TotalPasses
		for {
			record, err := csvdata.Reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				return err
			} else {
				if row == options.HeaderRowIndex {
					hrd = gd.FillHeaderRowStringData(row, pass, record)
				} else {

					// fill row data
					// it's a csv so we have string data
					rd := gd.FillRowStringData(row, pass, record)

					// get the row actions
					for _, rowaction := range rowparser.GetActions() {
						if hrd != nil {
							ra, ok := rowaction.(*gd.HeadedRowAction)
							if !ok {
								return errors.New("rowaction type was not a Headed Row Action")
							}
							ra.Header = hrd
							err := ra.Perform(rd)
							if err != nil {
								return err
							}
						} else {
							// perform action on the entire row data
							err := rowaction.Perform(rd)
							if err != nil {
								return err
							}
						}
					}
				}
			}
			row++
		}
		return nil

	}

	return errors.New("data source provided was not of type CSV File or was nil")

}
