package xlsx

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

	// standard xlsx row parser
	rowp := gd.NewRowProcessor(XLSXRowParse, DefaultXLXSAllSheetsProcessorOptions())
	gdp.AddProcessor(rowp)
	return gdp
}

// NewRowParser creates a Parser for a single file
func NewRowParserWithHeader(filepath string) *gd.Parser {
	gdp := gd.NewParser()
	file := gd.NewGridFile(filepath)
	gdp.AddSource(file)

	// standard xlsx row parser
	rowp := gd.NewHeaderRowProcessor(XLSXRowParse, DefaultXLXSAllSheetsHeaderProcessorOptions())
	gdp.AddProcessor(rowp)
	return gdp
}

func NewRowParserDefaultAction(filepath string, action acts.PerCell) *gd.Parser {
	return NewRowParserWithAction(filepath, DefaultXLXSAllSheetsProcessorOptions(), action)
}

// NewRowParser creates a Parser for a single file
func NewRowParserWithAction(filepath string, opts *XLXSOptions, action acts.PerCell) *gd.Parser {
	gdp := gd.NewParser()
	file := gd.NewGridFile(filepath)
	gdp.AddSource(file)

	// standard xlsx row parser
	rowp := gd.NewRowProcessor(XLSXRowParse, opts)
	rowp.AddAction(&action)

	// add processor
	gdp.AddProcessor(rowp)
	return gdp
}

func NewRowParserWithDefaultHeaderAction(filepath string, action acts.HeadedRow) *gd.Parser {
	return NewRowParserWithActionAndHeader(filepath, DefaultXLXSAllSheetsHeaderProcessorOptions(), action)
}

// NewRowParserWithActionAndHeader creates a Parser for a single file
func NewRowParserWithActionAndHeader(filepath string, opts *XLXSOptions, action acts.HeadedRow) *gd.Parser {
	gdp := gd.NewParser()
	file := gd.NewGridFile(filepath)
	gdp.AddSource(file)

	// standard xlsx row parser
	rowp := gd.NewHeaderRowProcessor(XLSXRowParse, opts)
	rowp.AddAction(&action)

	// add processor
	gdp.AddProcessor(rowp)
	return gdp
}

// NewHeadedRowParserWithAction creates a Parser for a single file
func NewHeadedRowParserWithAction(filepath string, action acts.HeadedRow, opts ...idata.SetOpt) *gd.Parser {
	gdp := gd.NewParser()

	file := gd.NewGridFile(filepath)
	// add the source
	gdp.AddSource(file)

	// set the options
	o := DefaultXLXSAllSheetsHeaderProcessorOptions()
	for _, setopt := range opts {
		setopt(o)
	}

	// standard csv row parser
	rowp := gd.NewHeaderRowProcessor(XLSXRowParse, o)

	// add action
	rowp.AddAction(&action)

	gdp.AddProcessor(rowp)
	return gdp
}
