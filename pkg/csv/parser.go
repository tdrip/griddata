package csv

import (
	gd "github.com/tdrip/griddata/pkg"
)

// CreateFileParser Creates a Parser for a single file
func CreateFileParser(filepath string) *gd.Parser {
	gdp := gd.CreateParser()
	file := CreateCSVFile(filepath)
	gdp.AddDataSource(file)

	// standard csv row parser
	rowp := gd.CreateRowProcessor(CSVParse)
	gdp.AddProcessor(rowp)
	return gdp
}

// CreateFileParser Creates a Parser for a single file
func CreateFileParserWithHeader(filepath string, headerowindex int) *gd.Parser {
	gdp := gd.CreateParser()
	file := CreateCSVFile(filepath)
	gdp.AddDataSource(file)

	// standard csv row parser
	rowp := gd.CreateHeaderRowProcessor(CSVParse, headerowindex)
	gdp.AddProcessor(rowp)
	return gdp
}

// CreateFileParser Creates a Parser for a single file
func CreateFileParserWithAction(filepath string, action *gd.RowAction) *gd.Parser {
	gdp := gd.CreateParser()
	file := CreateCSVFile(filepath)
	gdp.AddDataSource(file)

	// standard csv row parser
	rowp := gd.CreateRowProcessor(CSVParse)
	rowp.AddAction(action)
	gdp.AddProcessor(rowp)
	return gdp
}

// CreateFileParser Creates a Parser for a single file
func CreateFileParserWithActionAndHeader(filepath string, headerowindex int, action *gd.RowAction) *gd.Parser {
	gdp := gd.CreateParser()
	file := CreateCSVFile(filepath)
	gdp.AddDataSource(file)

	// standard csv row parser
	rowp := gd.CreateHeaderRowProcessor(CSVParse, headerowindex)
	rowp.AddAction(action)
	gdp.AddProcessor(rowp)
	return gdp
}
