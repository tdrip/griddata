package igrid

//IParser this is the parser interface
type IParser interface {
	GetRowParsers() []IRowParser
	SetRowParsers(rparsers []IRowParser)

	GetColumnParsers() []IColumnParser
	SetColumnParsers(cparsers []IColumnParser)

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
}
