package csv

import (
	gd "github.com/tdrip/griddata/pkg/data"
)

// NewRowParser creates a Parser for a single file
func NewRowParser(filepath string) *gd.Parser {
	gdp := gd.NewParser()
	file := gd.NewGridFile(filepath)
	gdp.AddSource(file)

	// standard csv row parser
	rowp := gd.NewRowProcessor(CSVRowParse, DefaultCSVOptions())
	gdp.AddProcessor(rowp)
	return gdp
}

// NewRowParser creates a Parser for a single file
func NewRowParserWithOptions(filepath string) *gd.Parser {
	gdp := gd.NewParser()
	file := gd.NewGridFile(filepath)
	gdp.AddSource(file)

	// standard csv row parser
	rowp := gd.NewHeaderRowProcessor(CSVRowParse, DefaultCSVHeaderOptions())
	gdp.AddProcessor(rowp)
	return gdp
}

func NewRowParserDefaultAction(filepath string, action gd.RowAction) *gd.Parser {
	return NewRowParserWithAction(filepath, DefaultCSVOptions(), action)
}

// NewRowParser creates a Parser for a single file
func NewRowParserWithAction(filepath string, opts *CSVOptions, action gd.RowAction) *gd.Parser {
	gdp := gd.NewParser()
	file := gd.NewGridFile(filepath)
	gdp.AddSource(file)

	// standard csv row parser
	rowp := gd.NewRowProcessor(CSVRowParse, opts)
	rowp.AddAction(&action)
	gdp.AddProcessor(rowp)
	return gdp
}

func NewRowParserWithDefaultHeaderAction(filepath string, action gd.HeadedRowAction) *gd.Parser {
	return NewRowParserWithHeaderAction(filepath, DefaultCSVHeaderOptions(), action)
}

// NewRowParser creates a Parser for a single file
func NewRowParserWithHeaderAction(filepath string, opts *CSVOptions, action gd.HeadedRowAction) *gd.Parser {
	gdp := gd.NewParser()
	file := gd.NewGridFile(filepath)
	gdp.AddSource(file)

	// standard csv row parser
	rowp := gd.NewHeaderRowProcessor(CSVRowParse, opts)
	rowp.AddAction(&action)
	gdp.AddProcessor(rowp)
	return gdp
}
