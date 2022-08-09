package grid

import (
	logr "github.com/sirupsen/logrus"
	igrid "github.com/tdrip/griddata/pkg/interfaces"
)

//Parser Grid data Parser with structs
type Parser struct {
	// inherit from the engine interface
	igrid.IParser

	// for logging
	Logger *logr.Logger

	//Row parsers
	RowParsers []igrid.IRowParser

	//Column parsers
	ColumnParsers []igrid.IColumnParser

	//DataSource
	DataSources []igrid.IDataSource
}

//CreateParser Creates a Parser
func CreateParser(logger *logr.Logger) *Parser {

	parser := Parser{}
	parser.Logger = logger
	parser.RowParsers = []igrid.IRowParser{}
	parser.ColumnParsers = []igrid.IColumnParser{}
	parser.DataSources = []igrid.IDataSource{}
	return &parser
}

//GetRowParsers Get the row parsers
func (gdp *Parser) GetRowParsers() []igrid.IRowParser {
	return gdp.RowParsers
}

//SetRowParsers Set the row parsers
func (gdp *Parser) SetRowParsers(rparsers []igrid.IRowParser) {
	gdp.RowParsers = rparsers
}

//AddRowParser Add a single row parser
func (gdp *Parser) AddRowParser(rparser igrid.IRowParser) {
	rparsers := gdp.RowParsers
	rparsers = append(rparsers, rparser)
	gdp.RowParsers = rparsers
}

//GetColumnParsers Get the column parsers
func (gdp *Parser) GetColumnParsers() []igrid.IColumnParser {
	return gdp.ColumnParsers
}

//SetColumnParsers Set the column parsers
func (gdp *Parser) SetColumnParsers(cparsers []igrid.IColumnParser) {
	gdp.ColumnParsers = cparsers
}

//AddColumnParser Add a single column parser
func (gdp *Parser) AddColumnParser(cparser igrid.IColumnParser) {
	cparsers := gdp.ColumnParsers
	cparsers = append(cparsers, cparser)
	gdp.ColumnParsers = cparsers
}

//GetDataSources Get the data sources
func (gdp *Parser) GetDataSources() []igrid.IDataSource {
	return gdp.DataSources
}

//SetDataSources Set the data sources
func (gdp *Parser) SetDataSources(datasources []igrid.IDataSource) {
	gdp.DataSources = datasources
}

//AddDataSource Set the data sources
func (gdp *Parser) AddDataSource(datasource igrid.IDataSource) {
	datasources := gdp.DataSources
	datasources = append(datasources, datasource)
	gdp.DataSources = datasources
}

//Execute - run the Column or Row Parsers
func (gdp *Parser) Execute() error {

	// Get the parsers
	cparser := gdp.GetColumnParsers()
	rparser := gdp.GetRowParsers()

	// get the data sources and validate each one
	datasources := gdp.GetDataSources()

	// Let's go through the data sources
	for d := 0; d < len(datasources); d++ {

		// validate the source
		parserr := datasources[d].Validate()
		if parserr != nil {
			// let's stop on validation
			return parserr
		}

		// open the data source (assuming stream based data)
		parserr = datasources[d].Open()
		if parserr != nil {
			return parserr
		}

		failed := false
		// let's walk through the column parsers and parse each one against the datasource
		for p := 0; p < len(cparser); p++ {
			parserr = cparser[p].Parse(gdp, datasources[d])
			if parserr != nil {
				break
			}
		}

		if failed {
			datasources[d].Close()
			return parserr
		}

		// let's walk through the row parsers and parse each one against the datasource
		for r := 0; r < len(rparser); r++ {
			parserr = rparser[r].Parse(gdp, datasources[d])
			if parserr != nil {
				return parserr
			}
		}

		// close the data source
		datasources[d].Close()

		if failed {
			return parserr
		}

	}

	return nil
}
