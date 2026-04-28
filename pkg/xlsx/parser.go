package xlsx

import (
	gd "github.com/tdrip/griddata/pkg"
)

// CreateRowParser Creates a Parser for a single file
func CreateRowParser(filepath string) *gd.Parser {
	gdp := gd.CreateParser()
	file := gd.CreateGridFile(filepath)
	gdp.AddDataSource(file)

	// create standard xlslx options
	opts := CreateXLXSProcessorOptions([]string{})
	opts.Defaults()

	// standard xlsx row parser
	rowp := gd.CreateRowProcessor(XLSXRowParse, opts)
	gdp.AddProcessor(rowp)
	return gdp
}

// CreateRowParser Creates a Parser for a single file
func CreateRowParserWithHeader(filepath string, headerowindex int) *gd.Parser {
	gdp := gd.CreateParser()
	file := gd.CreateGridFile(filepath)
	gdp.AddDataSource(file)

	// create standard xlslx options
	opts := CreateXLXSProcessorOptions([]string{})
	opts.Defaults()
	opts.HeaderRowindex = headerowindex

	// standard xlsx row parser
	rowp := gd.CreateHeaderRowProcessor(XLSXRowParse, opts)
	gdp.AddProcessor(rowp)
	return gdp
}

// CreateRowParser Creates a Parser for a single file
func CreateRowParserWithAction(filepath string, action gd.RowAction) *gd.Parser {
	gdp := gd.CreateParser()
	file := gd.CreateGridFile(filepath)
	gdp.AddDataSource(file)

	// create standard xlslx options
	opts := CreateXLXSProcessorOptions([]string{})
	opts.Defaults()

	// standard xlsx row parser
	rowp := gd.CreateRowProcessor(XLSXRowParse, opts)
	rowp.AddAction(&action)

	// add processor
	gdp.AddProcessor(rowp)
	return gdp
}

// CreateRowParser Creates a Parser for a single file
func CreateRowParserWithActionAndHeader(filepath string, headerowindex int, action gd.HeadedRowAction) *gd.Parser {
	gdp := gd.CreateParser()
	file := gd.CreateGridFile(filepath)
	gdp.AddDataSource(file)

	// create standard xlslx options
	opts := CreateXLXSProcessorOptions([]string{})
	opts.Defaults()
	opts.HeaderRowindex = headerowindex

	// standard xlsx row parser
	rowp := gd.CreateHeaderRowProcessor(XLSXRowParse, opts)
	rowp.AddAction(&action)

	// add processor
	gdp.AddProcessor(rowp)
	return gdp
}
