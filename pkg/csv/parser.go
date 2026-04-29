package csv

import (
	acts "github.com/tdrip/griddata/pkg/actions"
	gd "github.com/tdrip/griddata/pkg/data"
	idata "github.com/tdrip/griddata/pkg/data/interfaces"
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

func NewRowParserDefaultAction(filepath string, action acts.PerCell) *gd.Parser {
	return NewRowParserWithAction(filepath, DefaultCSVOptions(), action)
}

// NewRowParser creates a Parser for a single file
func NewRowParserWithAction(filepath string, opts *CSVOptions, action acts.PerCell) *gd.Parser {
	gdp := gd.NewParser()
	file := gd.NewGridFile(filepath)
	gdp.AddSource(file)

	// standard csv row parser
	rowp := gd.NewRowProcessor(CSVRowParse, opts)
	rowp.AddAction(&action)
	gdp.AddProcessor(rowp)
	return gdp
}

func NewRowParserWithDefaultHeaderAction(filepath string, action acts.HeadedRow) *gd.Parser {
	return NewRowParserWithHeaderAction(filepath, DefaultCSVHeaderOptions(), action)
}

// NewRowParser creates a Parser for a single file
func NewRowParserWithHeaderAction(filepath string, opts *CSVOptions, action acts.HeadedRow) *gd.Parser {
	gdp := gd.NewParser()
	file := gd.NewGridFile(filepath)
	gdp.AddSource(file)

	// standard csv row parser
	rowp := gd.NewHeaderRowProcessor(CSVRowParse, opts)
	rowp.AddAction(&action)
	gdp.AddProcessor(rowp)
	return gdp
}

/*
// NewHeadedRowParserWithAction creates a Parser for a single file
func NewHeadedRowParserWithAction(file *gd.GridFile, action iacts.Action, opts ...idata.SetOpt) *gd.Parser {
	gdp := gd.NewParser()

	// add the source
	gdp.AddSource(file)

	// set the options
	o := DefaultCSVHeaderOptions()
	for _, setopt := range opts {
		setopt(o)
	}

	// standard csv row parser
	rowp := gd.NewHeaderRowProcessor(CSVRowParse, o)

	// add action
	rowp.AddAction(action)

	gdp.AddProcessor(rowp)
	return gdp
}
*/

// NewHeadedRowParserWithAction creates a Parser for a single file
func NewHeadedRowParserWithAction(filepath string, action acts.HeadedRow, opts ...idata.SetOpt) *gd.Parser {
	gdp := gd.NewParser()

	file := gd.NewGridFile(filepath)
	// add the source
	gdp.AddSource(file)

	// set the options
	o := DefaultCSVHeaderOptions()
	for _, setopt := range opts {
		setopt(o)
	}

	// standard csv row parser
	rowp := gd.NewHeaderRowProcessor(CSVRowParse, o)

	// add action
	rowp.AddAction(&action)

	gdp.AddProcessor(rowp)
	return gdp
}
