package csv

import (
	gd "github.com/tdrip/griddata/pkg"
	sli "github.com/tdrip/logger/pkg/interfaces"
)

//CreateFileParser Creates a Parser for a single file
func CreateFileParser(slog sli.ISimpleLogger, filepath string) *gd.Parser {
	gdp := gd.CreateParser(slog)

	file := CreateCSVFile(filepath)
	gdp.AddDataSource(file)

	rowp := CreateRowParser()
	gdp.AddRowParser(rowp)
	return gdp
}
