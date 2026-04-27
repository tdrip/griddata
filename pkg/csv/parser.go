package csv

import (
	gd "github.com/tdrip/griddata/pkg"
)

// CreateRowParser Creates a Parser for a single file
func CreateRowParser(filepath string) *gd.Parser {
	gdp := gd.CreateParser()
	file := CreateCSVFile(filepath)
	gdp.AddDataSource(file)

	// standard csv row parser
	rowp := gd.CreateRowProcessor(CSVRowParse)
	gdp.AddProcessor(rowp)
	return gdp
}

// CreateRowParser Creates a Parser for a single file
func CreateRowParserWithHeader(filepath string, headerowindex int) *gd.Parser {
	gdp := gd.CreateParser()
	file := CreateCSVFile(filepath)
	gdp.AddDataSource(file)

	// standard csv row parser
	rowp := gd.CreateHeaderRowProcessor(CSVRowParse, headerowindex)
	gdp.AddProcessor(rowp)
	return gdp
}

// CreateRowParser Creates a Parser for a single file
func CreateRowParserWithAction(filepath string, action gd.RowAction) *gd.Parser {
	gdp := gd.CreateParser()
	file := CreateCSVFile(filepath)
	gdp.AddDataSource(file)

	// standard csv row parser
	rowp := gd.CreateRowProcessor(CSVRowParse)
	rowp.AddAction(&action)
	gdp.AddProcessor(rowp)
	return gdp
}

// CreateRowParser Creates a Parser for a single file
func CreateRowParserWithActionAndHeader(filepath string, headerowindex int, action gd.HeadedRowAction) *gd.Parser {
	gdp := gd.CreateParser()
	file := CreateCSVFile(filepath)
	gdp.AddDataSource(file)

	// standard csv row parser
	rowp := gd.CreateHeaderRowProcessor(CSVRowParse, headerowindex)
	rowp.AddAction(&action)
	gdp.AddProcessor(rowp)
	return gdp
}
