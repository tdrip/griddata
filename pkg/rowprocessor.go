package grid

import (
	idata "github.com/tdrip/griddata/pkg/data/interfaces"
)

type RowParse func(rp *RowProcessor, parent idata.IParser, data idata.IDataSource) error

// NewRowProcessor News the row parser
func NewRowProcessor(parse RowParse, opts idata.IDataProcessorOptions) *RowProcessor {
	return NewHeaderRowProcessor(parse, opts)
}

// NewRowProcessor News the row parser
func NewHeaderRowProcessor(parse RowParse, opts idata.IDataProcessorOptions) *RowProcessor {
	rp := &RowProcessor{}
	rp.SetOptions(opts)
	empty := make(map[string]idata.IDataAction)
	rp.Actions = empty
	rp.ParseFunc = parse
	return rp
}

// RowProcessor parses a csv row by row
type RowProcessor struct {

	// inherit from the row parser
	idata.IDataProcessor

	Options idata.IDataProcessorOptions

	Actions map[string]idata.IDataAction

	ParseFunc RowParse
}

func (rp *RowProcessor) Parse(parent idata.IParser, data idata.IDataSource) error {
	if rp.ParseFunc != nil {
		return rp.ParseFunc(rp, parent, data)
	}
	// do nothing
	return nil
}

// GetOptions Get the options for the row parser
func (rp *RowProcessor) GetOptions() idata.IDataProcessorOptions {
	return rp.Options
}

// SetOptions Set the options for the row parser
func (rp *RowProcessor) SetOptions(options idata.IDataProcessorOptions) {
	rp.Options = options
}

// actions for the row
func (rp *RowProcessor) GetActions() map[string]idata.IDataAction {
	return rp.Actions
}

func (rp *RowProcessor) SetActions(actions []idata.IDataAction) {
	data := rp.Actions
	for _, action := range actions {
		data[action.GetId()] = action
	}
	rp.Actions = data
}

func (rp *RowProcessor) AddAction(action idata.IDataAction) {
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
	empty := make(map[string]idata.IDataAction)
	rp.Actions = empty
}
