package igrid

//IRowProcessingOptions Represents the options for parsing a data row
type IRowProcessingOptions interface {
	Defaults()
	String() string
}

//IRowProcessor This interface parses a row in the grid
type IRowProcessor interface {

	// Options for parsing the row
	GetOptions() IRowProcessingOptions
	SetOptions(options IRowProcessingOptions)

	// Parse the row
	Parse(parent IParser, data IDataSource) error

	// actions for the row
	GetActions() []IRowAction
	SetActions(data map[string]IRowAction)
	AddAction(action IRowAction)
}
