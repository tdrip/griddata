package data

import (
	iact "github.com/tdrip/griddata/pkg/actions/interfaces"
	idata "github.com/tdrip/griddata/pkg/data/interfaces"
)

type ParseRow func(rp *RowProcessor, parent idata.Parser, data idata.Source) error

// RowProcessor parses a csv row by row
type RowProcessor struct {

	// inherit from the row parser
	idata.Processor

	Options idata.ProcessorOpts

	Actions map[string]iact.Action

	ParseFunc ParseRow
}

// NewRowProcessor News the row parser
func NewRowProcessor(parse ParseRow, opts idata.ProcessorOpts) *RowProcessor {
	return NewHeaderRowProcessor(parse, opts)
}

// NewRowProcessor News the row parser
func NewHeaderRowProcessor(parse ParseRow, opts idata.ProcessorOpts) *RowProcessor {
	rp := &RowProcessor{}
	rp.SetOptions(opts)
	empty := make(map[string]iact.Action)
	rp.Actions = empty
	rp.ParseFunc = parse
	return rp
}

func (rp *RowProcessor) Parse(parent idata.Parser, data idata.Source) error {
	if rp.ParseFunc != nil {
		return rp.ParseFunc(rp, parent, data)
	}
	// do nothing
	return nil
}

// GetOptions Get the options for the row parser
func (rp *RowProcessor) GetOptions() idata.ProcessorOpts {
	return rp.Options
}

// SetOptions Set the options for the row parser
func (rp *RowProcessor) SetOptions(options idata.ProcessorOpts) {
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
