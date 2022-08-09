package csv

import (
	logr "github.com/sirupsen/logrus"
	gd "github.com/tdrip/griddata/pkg"
)

//CreateFileParser Creates a Parser for a single file
func CreateFileParser(slog *logr.Logger, filepath string) *gd.Parser {
	gdp := gd.CreateParser(slog)
	file := CreateCSVFile(filepath)
	gdp.AddDataSource(file)

	rowp := CreateRowParser()
	gdp.AddRowParser(rowp)
	return gdp
}
