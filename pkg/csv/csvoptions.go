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

	rowPasses int

	headerRowindex int

	minColumns int

	seperator rune
}

// Constructors
func NewTSVOptions() *CSVOptions {
	opts := NewCSVOptions()
	sepf := WithSeperator(TSVSep)
	sepf(opts)
	return opts
}

func NewCSVOptions() *CSVOptions {
	opts := CSVOptions{}
	opts.Defaults()
	opts.seperator = CSVSep
	return &opts
}

// With Option Functions

func WithHeaderIndex(rowindex int) idata.SetOpt {
	return func(opts idata.ProcessorOpts) {
		xol := opts.(*CSVOptions)
		xol.headerRowindex = rowindex
	}
}

func WithMinColumns(numofcolumns int) idata.SetOpt {
	return func(opts idata.ProcessorOpts) {
		xol := opts.(*CSVOptions)
		xol.minColumns = numofcolumns
	}
}

func WithRowPasses(passes int) idata.SetOpt {
	return func(opts idata.ProcessorOpts) {
		xol := opts.(*CSVOptions)
		xol.rowPasses = passes
	}
}

func WithSeperator(seperator rune) idata.SetOpt {
	return func(opts idata.ProcessorOpts) {
		xol := opts.(*CSVOptions)
		xol.seperator = seperator
	}
}

// Defaults
func (rpo *CSVOptions) Defaults() {
	//Only pass over the row once
	rpo.rowPasses = idata.DEFAULTNUMOFPASSES

	// default to no header
	rpo.headerRowindex = idata.NOHEADERROWINDEX

	rpo.minColumns = idata.IGNORECOLUMNCOUNT
}

func (rpo *CSVOptions) RowPasses() int {
	return rpo.rowPasses
}

func (rpo *CSVOptions) HeaderRowIndex() int {
	return rpo.headerRowindex
}

func (rpo *CSVOptions) MinColumns() int {
	return rpo.minColumns
}

func (rpo *CSVOptions) Seperator() rune {
	return rpo.seperator
}

// String the readable version of the options
func (rpo *CSVOptions) String() string {
	return fmt.Sprintf("Total Row Passes: %d, Header Index: %d, Number of Columns Index: %d, Seperator: %s", rpo.RowPasses(), rpo.HeaderRowIndex(), rpo.MinColumns(), string(rpo.Seperator()))
}
