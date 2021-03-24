package igrid

//IRowParsingOptions Represents the options for parsing a data row
type IRowParsingOptions interface {
	Defaults()
	String() string
}

//IRowParser This interface parses a row in the grid
type IRowParser interface {

	// Options for parsing the row
	GetOptions() IRowParsingOptions
	SetOptions(options IRowParsingOptions)

	// Parse the row
	Parse(parent IParser, data IDataSource) error

	// actions for the row
	GetActions() []IRowAction
	SetActions(data map[string]IRowAction)
	AddAction(action IRowAction)
}
