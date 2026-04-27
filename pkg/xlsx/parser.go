package xlsx

import (
	gd "github.com/tdrip/griddata/pkg"
)

// CreateRowParser Creates a Parser for a single file
func CreateRowParser(filepath string) *gd.Parser {
	gdp := gd.CreateParser()
	file := CreateXLSXFile(filepath)
	gdp.AddDataSource(file)

	// standard xlsx row parser
	rowp := gd.CreateRowProcessor(XLSXRowParse)
	gdp.AddProcessor(rowp)
	return gdp
}

// CreateRowParser Creates a Parser for a single file
func CreateRowParserWithHeader(filepath string, headerowindex int) *gd.Parser {
	gdp := gd.CreateParser()
	file := CreateXLSXFile(filepath)
	gdp.AddDataSource(file)

	// standard xlsx row parser
	rowp := gd.CreateHeaderRowProcessor(XLSXRowParse, headerowindex)
	gdp.AddProcessor(rowp)
	return gdp
}

// CreateRowParser Creates a Parser for a single file
func CreateRowParserWithAction(filepath string, action gd.RowAction) *gd.Parser {
	gdp := gd.CreateParser()
	file := CreateXLSXFile(filepath)
	gdp.AddDataSource(file)

	// standard xlsx row parser
	rowp := gd.CreateRowProcessor(XLSXRowParse)
	rowp.AddAction(&action)
	gdp.AddProcessor(rowp)
	return gdp
}

// CreateRowParser Creates a Parser for a single file
func CreateRowParserWithActionAndHeader(filepath string, headerowindex int, action gd.HeadedRowAction) *gd.Parser {
	gdp := gd.CreateParser()
	file := CreateXLSXFile(filepath)
	gdp.AddDataSource(file)

	// standard xlsx row parser
	rowp := gd.CreateHeaderRowProcessor(XLSXRowParse, headerowindex)
	rowp.AddAction(&action)
	gdp.AddProcessor(rowp)
	return gdp
}
