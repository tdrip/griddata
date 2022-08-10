package igrid

//IRowProcessorOptions Represents the options for parsing a data row
type IRowProcessorOptions interface {
	Defaults()
	String() string
}

//IRowProcessor This interface parses a row in the grid
type IRowProcessor interface {

	// Options for parsing the row
	GetOptions() IRowProcessorOptions
	SetOptions(options IRowProcessorOptions)

	// Parse the row
	Parse(parent IParser, data IDataSource) error

	// actions for the row
	GetActions() map[string]IRowAction
	SetActions(actions []IRowAction)
	AddAction(action IRowAction)
	RemoveAction(string)
	ClearActions()
}
