package xlsx

import (
	"errors"
	"fmt"

	gd "github.com/tdrip/griddata/pkg"
	igrid "github.com/tdrip/griddata/pkg/interfaces"
	"github.com/xuri/excelize/v2"
)

// Parse parse the data source
func XLSXRowParse(rowparser *gd.RowProcessor, parent igrid.IParser, data igrid.IDataSource) error {
	// convert the idatasource to what we expect which is a XLSX File
	cdata, ok := data.(*XLSXFile)
	if !ok {
		return errors.New("data type was not a XLSX File")
	}
	// We need a GD Parser for the logging
	opts := rowparser.GetOptions()
	if opts == nil {
		return errors.New("options were nil")
	}
	options, ok := opts.(*gd.RowProcessorOptions)
	if !ok {
		return errors.New("options type was not a Row Processor Options")
	}

	hrd := &gd.RowData{}
	hrd = nil
	if cdata != nil {

		reader, err := excelize.OpenReader(cdata.Filestream)
		if err != nil {
			return err
		}
		pass := options.TotalPasses
		numcols := options.NumOfColumns
		sheets := cdata.Sheets
		if len(sheets) == 0 {
			sheets = reader.GetSheetList()
		}
		for _, sheet := range sheets {
			records, err := reader.GetRows(sheet)
			if err != nil {
				return err
			}
			for row, record := range records {
				if numcols > 0 {
					if len(record) != options.NumOfColumns {
						return fmt.Errorf("expected number of columns is %d but data source has %d", options.NumOfColumns, len(record))
					}
				}

				// are we dealing with a header row?
				if options.HeaderRowIndex >= 0 && row == options.HeaderRowIndex {
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
							ra, ok := rowaction.(*gd.HeadedRowAction)
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
		}

		return nil

	}

	return errors.New("data source provided was not of type CSV File or was nil")

}
