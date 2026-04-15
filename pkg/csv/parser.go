package csv

import (
	logr "github.com/sirupsen/logrus"
	gd "github.com/tdrip/griddata/pkg"
)

// CreateFileParser Creates a Parser for a single file
func CreateFileParser(slog *logr.Logger, filepath string) *gd.Parser {
	gdp := gd.CreateParser(slog)
	file := CreateCSVFile(filepath)
	gdp.AddDataSource(file)

	// standard csv row parser
	rowp := gd.CreateRowProcessor(CSVParse)
	gdp.AddProcessor(rowp)
	return gdp
}

// CreateFileParser Creates a Parser for a single file
func CreateFileParserWithHeader(slog *logr.Logger, filepath string, headerowindex int) *gd.Parser {
	gdp := gd.CreateParser(slog)
	file := CreateCSVFile(filepath)
	gdp.AddDataSource(file)

	// standard csv row parser
	rowp := gd.CreateHeaderRowProcessor(CSVParse, headerowindex)
	gdp.AddProcessor(rowp)
	return gdp
}

// CreateFileParser Creates a Parser for a single file
func CreateFileParserWithAction(slog *logr.Logger, filepath string, action *gd.RowAction) *gd.Parser {
	gdp := gd.CreateParser(slog)
	file := CreateCSVFile(filepath)
	gdp.AddDataSource(file)

	// standard csv row parser
	rowp := gd.CreateRowProcessor(CSVParse)
	rowp.AddAction(action)
	gdp.AddProcessor(rowp)
	return gdp
}

// CreateFileParser Creates a Parser for a single file
func CreateFileParserWithActionAndHeader(slog *logr.Logger, filepath string, headerowindex int, action *gd.RowAction) *gd.Parser {
	gdp := gd.CreateParser(slog)
	file := CreateCSVFile(filepath)
	gdp.AddDataSource(file)

	// standard csv row parser
	rowp := gd.CreateHeaderRowProcessor(CSVParse, headerowindex)
	rowp.AddAction(action)
	gdp.AddProcessor(rowp)
	return gdp
}
