package data

import (
	idata "github.com/tdrip/griddata/pkg/data/interfaces"
)

// Parser Grid data Parser with structs
type Parser struct {
	// inherit from the engine interface
	idata.Parser

	//Processors
	Processors []idata.Processor

	//Source
	Sources []idata.Source
}

// NewParser creates a Parser
func NewParser() *Parser {
	parser := Parser{
		Processors: []idata.Processor{},
		Sources:    []idata.Source{},
	}
	return &parser
}

// GetProcessors Get the processors
func (gdp *Parser) GetProcessors() []idata.Processor {
	return gdp.Processors
}

// SetRowProcessors Set the row processors
func (gdp *Parser) SetProcessors(rparsers []idata.Processor) {
	gdp.Processors = rparsers
}

// AddProcessor Add a single row processor
func (gdp *Parser) AddProcessor(rparser idata.Processor) {
	gdp.Processors = append(gdp.Processors, rparser)
}

// GetDataSources Get the data sources
func (gdp *Parser) GetSources() []idata.Source {
	return gdp.Sources
}

// SetDataSources Set the data sources
func (gdp *Parser) SetSources(datasources []idata.Source) {
	gdp.Sources = datasources
}

// AddDataSource Set the data sources
func (gdp *Parser) AddSource(source idata.Source) {
	gdp.Sources = append(gdp.Sources, source)
}

// Execute - run the Column or Row Parsers
func (gdp *Parser) Execute() error {

	// Get the processors
	processors := gdp.GetProcessors()

	// get the data sources and validate each one
	datasources := gdp.GetSources()

	// Let's go through the data sources an validate them first
	// don't open files or make a mess and leave
	for _, ds := range datasources {

		// validate the source
		err := ds.Validate()
		if err != nil {
			// let's stop on validation
			return err
		}

		defer ds.Close()

		// open the data source (assuming stream based data)
		err = ds.Open()
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
	datasources := gdp.GetSources()

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
