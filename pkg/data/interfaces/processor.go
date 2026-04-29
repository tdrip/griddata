package idata

import iaction "github.com/tdrip/griddata/pkg/actions/interfaces"

// Processor This interface parses a row in the grid
type Processor interface {

	// Options for parsing the row
	GetOptions() ProcessorOpts
	SetOptions(options ProcessorOpts)

	// Parse the row
	Parse(parent Parser, data Source) error

	// actions for the row or column
	GetActions() map[string]iaction.Action
	SetActions(actions []iaction.Action)
	AddAction(action iaction.Action)
	RemoveAction(string)
	ClearActions()
}
