package igrid

//IParser this is the parser interface
type IParser interface {
	GetRowParsers() []IRowProcessor
	SetRowParsers(rparsers []IRowProcessor)
	AddRowParser(rparser IRowProcessor)

	GetColumnParsers() []IColumnParser
	SetColumnParsers(cparsers []IColumnParser)
	AddColumnParser(cparser IColumnParser)

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
