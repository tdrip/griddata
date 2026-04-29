package xlsx

import (
	"fmt"
	"strings"

	idata "github.com/tdrip/griddata/pkg/data/interfaces"
)

// XLXSOptions number of passes etc
type XLXSOptions struct {

	// inherit from the row IRowProcessingOptions
	idata.ProcessorOpts

	rowPasses int

	headerRowindex int

	minColumns int

	sheets []string
}

func NewXLXSOptions() *XLXSOptions {
	opts := XLXSOptions{}
	opts.Defaults()
	return &opts
}

func WithHeaderIndex(rowindex int) idata.SetOpt {
	return func(opts idata.ProcessorOpts) {
		xol := opts.(*XLXSOptions)
		xol.headerRowindex = rowindex
	}
}

func WithSheetName(name string) idata.SetOpt {
	return func(opts idata.ProcessorOpts) {
		xol := opts.(*XLXSOptions)
		xol.sheets = append(xol.sheets, name)
	}
}

func WithMinColumns(numofcolumns int) idata.SetOpt {
	return func(opts idata.ProcessorOpts) {
		xol := opts.(*XLXSOptions)
		xol.minColumns = numofcolumns
	}
}

func WithRowPasses(passes int) idata.SetOpt {
	return func(opts idata.ProcessorOpts) {
		xol := opts.(*XLXSOptions)
		xol.rowPasses = passes
	}
}

func AllSheets(name string) idata.SetOpt {
	return func(opts idata.ProcessorOpts) {
		xol := opts.(*XLXSOptions)
		xol.sheets = []string{}
	}
}

// Defaults
func (rpo *XLXSOptions) Defaults() {
	//Only pass over the row once
	rpo.rowPasses = idata.DEFAULTNUMOFPASSES

	// default to no header
	rpo.headerRowindex = idata.NOHEADERROWINDEX

	rpo.minColumns = idata.IGNORECOLUMNCOUNT

	rpo.sheets = []string{}
}

func (rpo *XLXSOptions) RowPasses() int {
	return rpo.rowPasses
}

func (rpo *XLXSOptions) HeaderRowIndex() int {
	return rpo.headerRowindex
}

func (rpo *XLXSOptions) MinColumns() int {
	return rpo.minColumns
}

func (rpo *XLXSOptions) Sheets() []string {
	return rpo.sheets
}

// String the readable version of the options
func (rpo *XLXSOptions) String() string {
	return fmt.Sprintf("Total Row Passes: %d, Header Index: %d, Number of Columns Index: %d, Sheets: %s", rpo.RowPasses(), rpo.HeaderRowIndex(), rpo.MinColumns(), strings.Join(rpo.Sheets(), ","))
}
