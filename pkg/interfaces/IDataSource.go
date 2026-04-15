package igrid

// IDataSource The data source for the Parser
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
