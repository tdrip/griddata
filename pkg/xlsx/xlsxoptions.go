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

	Passes int `json:"total_passes"`

	HeaderRowindex int `json:"header_row_index"`

	NumOfcolumns int `json:"number_columns"`

	Sheets []string `json:"sheet_names"`
}

func DefaultXLXSAllSheetsHeaderProcessorOptions() *XLXSOptions {
	opts := NewXLXSOptions([]string{})
	opts.HeaderRowindex = 0
	return opts
}

func DefaultXLXSAllSheetsProcessorOptions() *XLXSOptions {
	return NewXLXSOptions([]string{})
}

func NewXLXSOptions(sheets []string) *XLXSOptions {
	opts := XLXSOptions{}
	opts.Defaults()
	opts.Sheets = sheets
	return &opts
}

func WithHeaderIndex(rowindex int) idata.SetOpt {
	return func(opts idata.ProcessorOpts) {
		xol := opts.(*XLXSOptions)
		xol.HeaderRowindex = rowindex
	}
}

func WithSheetName(name string) idata.SetOpt {
	return func(opts idata.ProcessorOpts) {
		xol := opts.(*XLXSOptions)
		xol.Sheets = append(xol.Sheets, name)
	}
}

func WithNumOfcolumns(numofcolumns int) idata.SetOpt {
	return func(opts idata.ProcessorOpts) {
		xol := opts.(*XLXSOptions)
		xol.NumOfcolumns = numofcolumns
	}
}

func WithPasses(passes int) idata.SetOpt {
	return func(opts idata.ProcessorOpts) {
		xol := opts.(*XLXSOptions)
		xol.Passes = passes
	}
}

func AllSheets(name string) idata.SetOpt {
	return func(opts idata.ProcessorOpts) {
		xol := opts.(*XLXSOptions)
		xol.Sheets = []string{}
	}
}

// Defaults
func (rpo *XLXSOptions) Defaults() {
	//Only pass over the row once
	rpo.Passes = idata.DEFAULTNUMOFPASSES

	// default to no header
	rpo.HeaderRowindex = idata.NOHEADERROWINDEX

	rpo.NumOfcolumns = idata.IGNORECOLUMNCOUNT

	rpo.Sheets = []string{}
}

func (rpo *XLXSOptions) TotalPasses() int {
	return rpo.Passes
}

func (rpo *XLXSOptions) HeaderRowIndex() int {
	return rpo.HeaderRowindex
}

func (rpo *XLXSOptions) NumOfColumns() int {
	return rpo.NumOfcolumns
}

// String the readable version of the options
func (rpo *XLXSOptions) String() string {
	return fmt.Sprintf("Total Row Passes: %d, Header Index: %d, Number of Columns Index: %d, Sheets: %s", rpo.TotalPasses(), rpo.HeaderRowIndex(), rpo.NumOfColumns(), strings.Join(rpo.Sheets, ","))
}
