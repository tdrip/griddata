package idata

// IDataProcessor This interface parses a row in the grid
type IDataProcessor interface {

	// Options for parsing the row
	GetOptions() IDataProcessorOptions
	SetOptions(options IDataProcessorOptions)

	// Parse the row
	Parse(parent IParser, data IDataSource) error

	// actions for the row or column
	GetActions() map[string]IDataAction
	SetActions(actions []IDataAction)
	AddAction(action IDataAction)
	RemoveAction(string)
	ClearActions()
}
