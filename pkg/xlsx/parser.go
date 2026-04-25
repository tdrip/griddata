package xlsx

import (
	gd "github.com/tdrip/griddata/pkg"
)

// CreateFileParser Creates a Parser for a single file
func CreateFileParser(filepath string) *gd.Parser {
	gdp := gd.CreateParser()
	file := CreateXLSXFile(filepath)
	gdp.AddDataSource(file)

	// standard xlsx row parser
	rowp := gd.CreateRowProcessor(XLSXRowParse)
	gdp.AddProcessor(rowp)
	return gdp
}

// CreateFileParser Creates a Parser for a single file
func CreateFileParserWithHeader(filepath string, headerowindex int) *gd.Parser {
	gdp := gd.CreateParser()
	file := CreateXLSXFile(filepath)
	gdp.AddDataSource(file)

	// standard xlsx row parser
	rowp := gd.CreateHeaderRowProcessor(XLSXRowParse, headerowindex)
	gdp.AddProcessor(rowp)
	return gdp
}

// CreateFileParser Creates a Parser for a single file
func CreateFileParserWithAction(filepath string, action gd.RowAction) *gd.Parser {
	gdp := gd.CreateParser()
	file := CreateXLSXFile(filepath)
	gdp.AddDataSource(file)

	// standard xlsx row parser
	rowp := gd.CreateRowProcessor(XLSXRowParse)
	rowp.AddAction(&action)
	gdp.AddProcessor(rowp)
	return gdp
}

// CreateFileParser Creates a Parser for a single file
func CreateFileParserWithActionAndHeader(filepath string, headerowindex int, action gd.HeadedRowAction) *gd.Parser {
	gdp := gd.CreateParser()
	file := CreateXLSXFile(filepath)
	gdp.AddDataSource(file)

	// standard xlsx row parser
	rowp := gd.CreateHeaderRowProcessor(XLSXRowParse, headerowindex)
	rowp.AddAction(&action)
	gdp.AddProcessor(rowp)
	return gdp
}
