package grid

import (
	igrid "github.com/tdrip/griddata/pkg/interfaces"
)

// Parser Grid data Parser with structs
type Parser struct {
	// inherit from the engine interface
	igrid.IParser

	//Processors
	Processors []igrid.IDataProcessor

	//DataSource
	DataSources []igrid.IDataSource
}

// CreateParser Creates a Parser
func CreateParser() *Parser {
	parser := Parser{
		Processors:  []igrid.IDataProcessor{},
		DataSources: []igrid.IDataSource{},
	}
	return &parser
}

// GetProcessors Get the processors
func (gdp *Parser) GetProcessors() []igrid.IDataProcessor {
	return gdp.Processors
}

// SetRowProcessors Set the row processors
func (gdp *Parser) SetProcessors(rparsers []igrid.IDataProcessor) {
	gdp.Processors = rparsers
}

// AddProcessor Add a single row processor
func (gdp *Parser) AddProcessor(rparser igrid.IDataProcessor) {
	gdp.Processors = append(gdp.Processors, rparser)
}

// GetDataSources Get the data sources
func (gdp *Parser) GetDataSources() []igrid.IDataSource {
	return gdp.DataSources
}

// SetDataSources Set the data sources
func (gdp *Parser) SetDataSources(datasources []igrid.IDataSource) {
	gdp.DataSources = datasources
}

// AddDataSource Set the data sources
func (gdp *Parser) AddDataSource(datasource igrid.IDataSource) {
	gdp.DataSources = append(gdp.DataSources, datasource)
}

// Execute - run the Column or Row Parsers
func (gdp *Parser) Execute() error {

	// Get the processors
	processors := gdp.GetProcessors()

	// get the data sources and validate each one
	datasources := gdp.GetDataSources()

	// Let's go through the data sources an validate them first
	// don't open files or make a mess and leave
	for _, ds := range datasources {

		// validate the source
		err := ds.Validate()
		if err != nil {
			// let's stop on validation
			return err
		}
	}

	// now open files
	for _, ds := range datasources {

		defer ds.Close()

		// open the data source (assuming stream based data)
		err := ds.Open()
		if err != nil {
			return err
		}

		failed := false
		// let's walk through the column parsers and parse each one against the datasource
		for p := 0; p < len(processors); p++ {
			err = processors[p].Parse(gdp, ds)
			if err != nil {
				failed = true
				break
			}
		}

		if failed {
			return err
		}

	}

	return nil
}

// Close all data sources
func (gdp *Parser) Close() error {

	// get the data sources and validate each one
	datasources := gdp.GetDataSources()

	// Let's go through the data sources
	for _, ds := range datasources {
		// close the data source
		err := ds.Close()
		if err != nil {
			return err
		}

	}

	return nil
}
