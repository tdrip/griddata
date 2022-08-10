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

	//Processors
	Processors []igrid.IDataProcessor

	//DataSource
	DataSources []igrid.IDataSource
}

//CreateParser Creates a Parser
func CreateParser(logger *logr.Logger) *Parser {

	parser := Parser{}
	parser.Logger = logger
	parser.Processors = []igrid.IDataProcessor{}
	parser.DataSources = []igrid.IDataSource{}
	return &parser
}

//GetProcessors Get the processors
func (gdp *Parser) GetProcessors() []igrid.IDataProcessor {
	return gdp.Processors
}

//SetRowProcessors Set the row processors
func (gdp *Parser) SetProcessors(rparsers []igrid.IDataProcessor) {
	gdp.Processors = rparsers
}

//AddProcessor Add a single row processor
func (gdp *Parser) AddProcessor(rparser igrid.IDataProcessor) {
	rparsers := gdp.Processors
	rparsers = append(rparsers, rparser)
	gdp.Processors = rparsers
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

	// Get the processors
	processors := gdp.GetProcessors()

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
		for p := 0; p < len(processors); p++ {
			parserr = processors[p].Parse(gdp, datasources[d])
			if parserr != nil {
				failed = true
				break
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
