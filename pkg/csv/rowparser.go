package csv

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"

	act "github.com/tdrip/griddata/pkg/actions"
	gd "github.com/tdrip/griddata/pkg/data"
	idata "github.com/tdrip/griddata/pkg/data/interfaces"
)

// Parse parse the data source
func CSVRowParse(rowparser *gd.RowProcessor, parent idata.IParser, data idata.Source) error {

	// convert the idatasource to what we expect which is a CSV File
	cdata, ok := data.(*gd.GridFile)
	if !ok {
		return errors.New("data type was not a File")
	}
	// We need a GD Parser for the logging
	opts := rowparser.GetOptions()
	if opts == nil {
		return errors.New("row processor options were nil")
	}
	options, ok := opts.(*CSVOptions)
	if !ok {
		return errors.New("options type was not a Row Processor Options")
	}

	hrd := &gd.RowData{}
	hrd = nil
	if cdata != nil {
		row := 0
		pass := options.TotalPasses()
		numcols := options.NumOfColumns()
		if cdata.Filestream == nil {
			return errors.New("Filestream was nil - was it opened correctly")
		}
		reader := csv.NewReader(cdata.Filestream)
		if reader == nil {
			return errors.New("failed to create csv reader for file stream")
		}
		reader.Comma = options.Seperator
		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				return err
			} else {

				if numcols > 0 {
					if len(record) != numcols {
						return fmt.Errorf("expected number of columns is %d but data source has %d", numcols, len(record))
					}
				}

				// are we dealing with a header row?
				if options.HeaderRowIndex() >= 0 && row == options.HeaderRowIndex() {
					// set the header row data
					hrd = gd.FillRowStringData(row, pass, record)
				} else {
					// we have header row data
					if hrd != nil {
						// get the row actions
						for _, rowaction := range rowparser.GetActions() {
							rd, err := gd.FillHeaderRowStringData(row, pass, record, hrd)
							if err != nil {
								return err
							}
							ra, ok := rowaction.(*act.HeadedRow)
							if !ok {
								return errors.New("row action type was not a Headed Row Action")
							}
							err = ra.Perform(rd)
							if err != nil {
								return err
							}
						}
					} else {
						// fill row data but no header
						rd := gd.FillRowStringData(row, pass, record)

						// get the row actions
						for _, rowaction := range rowparser.GetActions() {

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
