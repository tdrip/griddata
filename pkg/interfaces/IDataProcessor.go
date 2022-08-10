package igrid

type IDataAction interface {
	GetId() string
	PeformAction(interface{})
}

//IDataProcessorOptions Represents the options for parsing a data row
type IDataProcessorOptions interface {
	Defaults()
	String() string
}

//IDataProcessor This interface parses a row in the grid
type IDataProcessor interface {

	// Options for parsing the row
	GetOptions() IDataProcessorOptions
	SetOptions(options IDataProcessorOptions)

	// Parse the row
	Parse(parent IParser, data IDataSource) error

	// actions for the row
	GetActions() map[string]IDataAction
	SetActions(actions []IDataAction)
	AddAction(action IDataAction)
	RemoveAction(string)
	ClearActions()
}
