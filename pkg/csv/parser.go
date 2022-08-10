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

	rowp := CreateRowProcessor()
	gdp.AddProcessor(rowp)
	return gdp
}

//CreateFileParser Creates a Parser for a single file
func CreateFileParserWithAction(slog *logr.Logger, filepath string, action *CSVRowAction) *gd.Parser {
	gdp := gd.CreateParser(slog)
	file := CreateCSVFile(filepath)
	gdp.AddDataSource(file)

	rowp := CreateRowProcessor()
	rowp.AddAction(action)
	gdp.AddProcessor(rowp)
	return gdp
}
