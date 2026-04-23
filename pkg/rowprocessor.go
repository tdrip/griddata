package grid

import (
	igrid "github.com/tdrip/griddata/pkg/interfaces"
)

type RowParse func(rp *RowProcessor, parent igrid.IParser, data igrid.IDataSource) error

// CreateRowProcessor creates the row parser
func CreateRowProcessor(parse RowParse) *RowProcessor {
	return CreateHeaderRowProcessor(parse, -1)
}

// CreateRowProcessor creates the row parser
func CreateHeaderRowProcessor(parse RowParse, headerowindex int) *RowProcessor {
	rp := &RowProcessor{}
	//Option
	opts := &RowProcessorOptions{}
	opts.Defaults()
	opts.HeaderRowIndex = headerowindex
	rp.SetOptions(opts)
	empty := make(map[string]igrid.IDataAction)
	rp.Actions = empty
	rp.ParseFunc = parse
	return rp
}

// RowProcessor parses a csv row by row
type RowProcessor struct {

	// inherit from the row parser
	igrid.IDataProcessor

	Options *RowProcessorOptions

	Actions map[string]igrid.IDataAction

	ParseFunc RowParse
}

func (rp *RowProcessor) Parse(parent igrid.IParser, data igrid.IDataSource) error {
	if rp.ParseFunc != nil {
		return rp.ParseFunc(rp, parent, data)
	}
	// do nothing
	return nil
}

// GetOptions Get the options for the row parser
func (rp *RowProcessor) GetOptions() igrid.IDataProcessorOptions {
	return rp.Options
}

// SetOptions Set the options for the row parser
func (rp *RowProcessor) SetOptions(options igrid.IDataProcessorOptions) {
	rp.Options = options.(*RowProcessorOptions)
}

// actions for the row
func (rp *RowProcessor) GetActions() map[string]igrid.IDataAction {
	return rp.Actions
}

func (rp *RowProcessor) SetActions(actions []igrid.IDataAction) {
	for _, action := range actions {
		rp.AddAction(action)
	}
}

func (rp *RowProcessor) AddAction(action igrid.IDataAction) {
	data := rp.Actions
	data[action.GetId()] = action
	rp.Actions = data
}

func (rp *RowProcessor) RemoveAction(id string) {
	data := rp.Actions
	delete(data, id)
	rp.Actions = data
}

func (rp *RowProcessor) ClearActions() {
	empty := make(map[string]igrid.IDataAction)
	rp.Actions = empty
}
