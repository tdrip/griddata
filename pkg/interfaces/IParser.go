package igrid

//IParser this is the parser interface
type IParser interface {
	GetRowProcessors() []IDataProcessor
	SetRowProcessors(rparsers []IDataProcessor)
	AddRowProcessor(rparser IDataProcessor)

	GetColumnParsers() []IDataProcessor
	SetColumnParsers(cparsers []IDataProcessor)
	AddColumnParser(cparser IDataProcessor)

	GetDataSources() []IDataSource
	SetDataSources(datasources []IDataSource)

	AddDataSource(datasources IDataSource)

	Execute() error
}

//IDataSource The data source for the Parser
type IDataSource interface {
	// Validate the source
	// for checks etc
	Validate() error

	Open() error

	Close() error
	/*
		Later
		// add child sources
		GetChildSources() []IDataSource
		SetChildSources([]IDataSource)
		AetChildSource(csource IDataSource)
	*/
}
