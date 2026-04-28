package idata

// Source The data source for the Parser
type Source interface {
	// Validate the source
	// for checks etc
	Validate() error

	Open() error

	Close() error
}
