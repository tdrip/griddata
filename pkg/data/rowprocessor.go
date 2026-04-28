package data

import (
	iact "github.com/tdrip/griddata/pkg/actions/interfaces"
	idata "github.com/tdrip/griddata/pkg/data/interfaces"
)

type RowParse func(rp *RowProcessor, parent idata.IParser, data idata.Source) error

// NewRowProcessor News the row parser
func NewRowProcessor(parse RowParse, opts idata.ProcessorOptions) *RowProcessor {
	return NewHeaderRowProcessor(parse, opts)
}

// NewRowProcessor News the row parser
func NewHeaderRowProcessor(parse RowParse, opts idata.ProcessorOptions) *RowProcessor {
	rp := &RowProcessor{}
	rp.SetOptions(opts)
	empty := make(map[string]iact.Action)
	rp.Actions = empty
	rp.ParseFunc = parse
	return rp
}

// RowProcessor parses a csv row by row
type RowProcessor struct {

	// inherit from the row parser
	idata.Processor

	Options idata.ProcessorOptions

	Actions map[string]iact.Action

	ParseFunc RowParse
}

func (rp *RowProcessor) Parse(parent idata.IParser, data idata.Source) error {
	if rp.ParseFunc != nil {
		return rp.ParseFunc(rp, parent, data)
	}
	// do nothing
	return nil
}

// GetOptions Get the options for the row parser
func (rp *RowProcessor) GetOptions() idata.ProcessorOptions {
	return rp.Options
}

// SetOptions Set the options for the row parser
func (rp *RowProcessor) SetOptions(options idata.ProcessorOptions) {
	rp.Options = options
}

// actions for the row
func (rp *RowProcessor) GetActions() map[string]iact.Action {
	return rp.Actions
}

func (rp *RowProcessor) SetActions(actions []iact.Action) {
	data := rp.Actions
	for _, action := range actions {
		data[action.GetId()] = action
	}
	rp.Actions = data
}

func (rp *RowProcessor) AddAction(action iact.Action) {
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
	empty := make(map[string]iact.Action)
	rp.Actions = empty
}
