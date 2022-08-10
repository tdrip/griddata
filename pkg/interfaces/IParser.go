package igrid

//IParser this is the parser interface
type IParser interface {
	GetProcessors() []IDataProcessor
	SetProcessors([]IDataProcessor)
	AddProcessor(IDataProcessor)

	GetDataSources() []IDataSource
	SetDataSources([]IDataSource)

	AddDataSource(IDataSource)

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
