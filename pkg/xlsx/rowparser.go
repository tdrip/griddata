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
	xsldata, ok := data.(*XLSXFile)
	if !ok {
		return errors.New("data type was not a XLSX File")
	}
	// We need a GD Parser for the logging
	opts := rowparser.GetOptions()
	options, ok := opts.(*gd.RowProcessorOptions)
	if !ok {
		return errors.New("options type was not a Row Processor Options")
	}

	var hrd *gd.RowData
	hrd = nil
	if xsldata != nil {

		reader, err := excelize.OpenReader(xsldata.Filestream)
		if err != nil {
			return err
		}
		pass := options.TotalPasses
		numcols := options.NumOfColumns
		sheets := xsldata.Sheets
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
				if row == options.HeaderRowIndex {
					hrd = gd.FillRowStringData(row, pass, record)
				} else {

					if hrd != nil {
						// get the row actions
						for _, rowaction := range rowparser.GetActions() {
							rd, err := gd.FillHeaderRowStringData(row, pass, record, hrd)
							if err != nil {
								return err
							}
							ra, ok := rowaction.(*gd.HeadedRowAction)
							if !ok {
								return errors.New("rowaction type was not a Headed Row Action")
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
