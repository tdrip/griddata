package xlsx

import (
	gd "github.com/tdrip/griddata/pkg"
)

// NewRowParser creates a Parser for a single file
func NewRowParser(filepath string) *gd.Parser {
	gdp := gd.NewParser()
	file := gd.NewGridFile(filepath)
	gdp.AddDataSource(file)

	// standard xlsx row parser
	rowp := gd.NewRowProcessor(XLSXRowParse, DefaultXLXSAllSheetsProcessorOptions())
	gdp.AddProcessor(rowp)
	return gdp
}

// NewRowParser creates a Parser for a single file
func NewRowParserWithHeader(filepath string) *gd.Parser {
	gdp := gd.NewParser()
	file := gd.NewGridFile(filepath)
	gdp.AddDataSource(file)

	// standard xlsx row parser
	rowp := gd.NewHeaderRowProcessor(XLSXRowParse, DefaultXLXSAllSheetsHeaderProcessorOptions())
	gdp.AddProcessor(rowp)
	return gdp
}

// NewRowParser creates a Parser for a single file
func NewRowParserWithAction(filepath string, opts *XLXSOptions, action gd.RowAction) *gd.Parser {
	gdp := gd.NewParser()
	file := gd.NewGridFile(filepath)
	gdp.AddDataSource(file)

	// standard xlsx row parser
	rowp := gd.NewRowProcessor(XLSXRowParse, opts)
	rowp.AddAction(&action)

	// add processor
	gdp.AddProcessor(rowp)
	return gdp
}

func NewRowParserWithDefaultHeaderAction(filepath string, action gd.HeadedRowAction) *gd.Parser {
	return NewRowParserWithActionAndHeader(filepath, DefaultXLXSAllSheetsHeaderProcessorOptions(), action)
}

// NewRowParserWithActionAndHeader creates a Parser for a single file
func NewRowParserWithActionAndHeader(filepath string, opts *XLXSOptions, action gd.HeadedRowAction) *gd.Parser {
	gdp := gd.NewParser()
	file := gd.NewGridFile(filepath)
	gdp.AddDataSource(file)

	// standard xlsx row parser
	rowp := gd.NewHeaderRowProcessor(XLSXRowParse, opts)
	rowp.AddAction(&action)

	// add processor
	gdp.AddProcessor(rowp)
	return gdp
}
