package csv

import (
	"fmt"

	idata "github.com/tdrip/griddata/pkg/data/interfaces"
)

const (
	TSVSep = '	'
	CSVSep = ','
)

// CSVOptions number of passes etc
type CSVOptions struct {

	// inherit from the row IRowProcessingOptions
	idata.ProcessorOpts

	Passes int `json:"passes"`

	HeaderRowindex int `json:"header_row_index"`

	NumOfcolumns int `json:"number_columns"`

	Seperator rune `json:"sheet_names"`
}

func DefaultCSVOptions() *CSVOptions {
	return NewCSVOptions(CSVSep)
}

func DefaultCSVHeaderOptions() *CSVOptions {
	opts := NewCSVOptions(CSVSep)
	opts.HeaderRowindex = 0
	return opts
}

func DefaultTSVOptions() *CSVOptions {
	opts := NewCSVOptions(TSVSep)
	return opts
}

func DefaultTSVHeaderOptions() *CSVOptions {
	opts := NewCSVOptions(TSVSep)
	opts.HeaderRowindex = 0
	return opts
}

func NewCSVOptions(seperator rune) *CSVOptions {
	opts := CSVOptions{}
	opts.Defaults()
	opts.Seperator = seperator
	return &opts
}

func WithHeaderIndex(rowindex int) idata.SetOpt {
	return func(opts idata.ProcessorOpts) {
		xol := opts.(*CSVOptions)
		xol.HeaderRowindex = rowindex
	}
}

func WithNumOfcolumns(numofcolumns int) idata.SetOpt {
	return func(opts idata.ProcessorOpts) {
		xol := opts.(*CSVOptions)
		xol.NumOfcolumns = numofcolumns
	}
}

func WithPasses(passes int) idata.SetOpt {
	return func(opts idata.ProcessorOpts) {
		xol := opts.(*CSVOptions)
		xol.Passes = passes
	}
}

func WithSeperator(seperator rune) idata.SetOpt {
	return func(opts idata.ProcessorOpts) {
		xol := opts.(*CSVOptions)
		xol.Seperator = seperator
	}
}

// Defaults
func (rpo *CSVOptions) Defaults() {
	//Only pass over the row once
	rpo.Passes = idata.DEFAULTNUMOFPASSES

	// default to no header
	rpo.HeaderRowindex = idata.NOHEADERROWINDEX

	rpo.NumOfcolumns = idata.IGNORECOLUMNCOUNT
}

func (rpo *CSVOptions) TotalPasses() int {
	return rpo.Passes
}

func (rpo *CSVOptions) HeaderRowIndex() int {
	return rpo.HeaderRowindex
}

func (rpo *CSVOptions) NumOfColumns() int {
	return rpo.NumOfcolumns
}

// String the readable version of the options
func (rpo *CSVOptions) String() string {
	return fmt.Sprintf("Total Row Passes: %d, Header Index: %d, Number of Columns Index: %d, Seperator: %s", rpo.TotalPasses(), rpo.HeaderRowIndex(), rpo.NumOfColumns(), string(rpo.Seperator))
}
