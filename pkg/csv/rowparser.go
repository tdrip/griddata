package csv

import (
	"errors"
	"io"

	logr "github.com/sirupsen/logrus"
	gd "github.com/tdrip/griddata/pkg"
	igrid "github.com/tdrip/griddata/pkg/interfaces"
)

// Parse parse the data source
func CSVParse(rowparser *gd.RowProcessor, parent igrid.IParser, data igrid.IDataSource) error {

	// convert the idatasource to what we expect which is a CSV File
	csvdata := data.(*CSVFile)

	// We need a GD Parser for the logging
	gdp := parent.(*gd.Parser)
	var slog *logr.Logger
	slog = nil
	if gdp.Logger != nil {
		slog = gdp.Logger.(*logr.Logger)
	}
	// We need a GD Parser for the logging
	opts := rowparser.GetOptions()
	options := opts.(*gd.RowProcessorOptions)
	var hrd *gd.RowData
	hrd = nil
	if csvdata != nil {
		row := 0
		pass := options.TotalPasses
		for {
			record, err := csvdata.Reader.Read()
			if err == io.EOF {
				if slog != nil {
					slog.Debug("csv parse - End of file")
				}
				break
			}
			if err != nil {
				if slog != nil {
					slog.Errorf("csv parse - %v", err)
				}
				return err
			} else {
				if slog != nil {
					slog.Debugf("csv parse - Record: %v", record)
				}
				if row == options.HeaderRowIndex {
					if slog != nil {
						slog.Debug("csv parse - Header row index")
					}
					hrd = CreateRowData(row, pass, record)
				} else {

					//create row data
					rd := CreateRowData(row, pass, record)

					// get the row actions
					for _, rowaction := range rowparser.GetActions() {
						if hrd != nil {
							ra := rowaction.(*gd.HeadedRowAction)
							if !ra.HasHeader() {
								ra.SetHeader(hrd)
							}
							err := ra.Peform(rd)
							if err != nil {
								return err
							}
						} else {
							// perform action on the entire row data
							err := rowaction.Peform(rd)
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

// CreateRowData creates a row data from a parsed CSV
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
