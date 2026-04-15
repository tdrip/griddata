package igrid

// IParser this is the parser interface
type IParser interface {
	GetProcessors() []IDataProcessor
	SetProcessors([]IDataProcessor)
	AddProcessor(IDataProcessor)

	GetDataSources() []IDataSource
	SetDataSources([]IDataSource)

	AddDataSource(IDataSource)

	Execute() error

	Close() error
}
