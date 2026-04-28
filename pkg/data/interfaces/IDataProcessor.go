package idata

// Processor This interface parses a row in the grid
type Processor interface {

	// Options for parsing the row
	GetOptions() ProcessorOptions
	SetOptions(options ProcessorOptions)

	// Parse the row
	Parse(parent IParser, data Source) error

	// actions for the row or column
	GetActions() map[string]Action
	SetActions(actions []Action)
	AddAction(action Action)
	RemoveAction(string)
	ClearActions()
}
