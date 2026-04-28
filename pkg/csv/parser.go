package csv

import (
	gd "github.com/tdrip/griddata/pkg"
)

// CreateRowParser Creates a Parser for a single file
func CreateRowParser(filepath string) *gd.Parser {
	gdp := gd.CreateParser()
	file := gd.CreateGridFile(filepath)
	gdp.AddDataSource(file)

	// create standard csv options
	opts := CreateCSVProcessorOptions(',')
	opts.Defaults()

	// standard csv row parser
	rowp := gd.CreateRowProcessor(CSVRowParse, opts)
	gdp.AddProcessor(rowp)
	return gdp
}

// CreateRowParser Creates a Parser for a single file
func CreateRowParserWithOptions(filepath string, opts *CSVProcessorOptions) *gd.Parser {
	gdp := gd.CreateParser()
	file := gd.CreateGridFile(filepath)
	gdp.AddDataSource(file)

	// standard csv row parser
	rowp := gd.CreateHeaderRowProcessor(CSVRowParse, opts)
	gdp.AddProcessor(rowp)
	return gdp
}

// CreateRowParser Creates a Parser for a single file
func CreateRowParserWithAction(filepath string, opts *CSVProcessorOptions, action gd.RowAction) *gd.Parser {
	gdp := gd.CreateParser()
	file := gd.CreateGridFile(filepath)
	gdp.AddDataSource(file)

	// standard csv row parser
	rowp := gd.CreateRowProcessor(CSVRowParse, opts)
	rowp.AddAction(&action)
	gdp.AddProcessor(rowp)
	return gdp
}

// CreateRowParser Creates a Parser for a single file
func CreateRowParserWithHeaderAction(filepath string, opts *CSVProcessorOptions, action gd.HeadedRowAction) *gd.Parser {
	gdp := gd.CreateParser()
	file := gd.CreateGridFile(filepath)
	gdp.AddDataSource(file)

	// standard csv row parser
	rowp := gd.CreateHeaderRowProcessor(CSVRowParse, opts)
	rowp.AddAction(&action)
	gdp.AddProcessor(rowp)
	return gdp
}
