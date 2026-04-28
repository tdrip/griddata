package csv

import (
	"fmt"

	igrid "github.com/tdrip/griddata/pkg/interfaces"
)

// CSVProcessorOptions number of passes etc
type CSVProcessorOptions struct {

	// inherit from the row IRowProcessingOptions
	igrid.IDataProcessorOptions

	Passes int `json:"total_passes"`

	HeaderRowindex int `json:"header_row_index"`

	NumOfcolumns int `json:"number_columns"`

	Seperator rune `json:"sheet_names"`
}

func CreateCSVProcessorOptions(seperator rune) *CSVProcessorOptions {
	opts := CSVProcessorOptions{}
	opts.Defaults()
	opts.Seperator = seperator
	return &opts
}

// Defaults
func (rpo *CSVProcessorOptions) Defaults() {
	//Only pass over the row once
	rpo.Passes = igrid.DEFAULTNUMOFPASSES

	// default to no header
	rpo.HeaderRowindex = igrid.NOHEADERROWINDEX

	rpo.NumOfcolumns = igrid.IGNORECOLUMNCOUNT
}

func (rpo *CSVProcessorOptions) TotalPasses() int {
	return rpo.Passes
}

func (rpo *CSVProcessorOptions) HeaderRowIndex() int {
	return rpo.HeaderRowindex
}

func (rpo *CSVProcessorOptions) NumOfColumns() int {
	return rpo.NumOfcolumns
}

// String the readable version of the options
func (rpo *CSVProcessorOptions) String() string {
	return fmt.Sprintf("Total Row Passes: %d, Header Index: %d, Number of Columns Index: %d, Seperator: %s", rpo.TotalPasses(), rpo.HeaderRowIndex(), rpo.NumOfColumns(), string(rpo.Seperator))
}
