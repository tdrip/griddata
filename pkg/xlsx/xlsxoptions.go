package xlsx

import (
	"fmt"
	"strings"

	igrid "github.com/tdrip/griddata/pkg/interfaces"
)

// XLXSProcessorOptions number of passes etc
type XLXSProcessorOptions struct {

	// inherit from the row IRowProcessingOptions
	igrid.IDataProcessorOptions

	Passes int `json:"total_passes"`

	HeaderRowindex int `json:"header_row_index"`

	NumOfcolumns int `json:"number_columns"`

	Sheets []string `json:"sheet_names"`
}

func CreateXLXSProcessorOptions(sheets []string) *XLXSProcessorOptions {
	opts := XLXSProcessorOptions{}
	opts.Defaults()
	opts.Sheets = sheets
	return &opts
}

// Defaults
func (rpo *XLXSProcessorOptions) Defaults() {
	//Only pass over the row once
	rpo.Passes = igrid.DEFAULTNUMOFPASSES

	// default to no header
	rpo.HeaderRowindex = igrid.NOHEADERROWINDEX

	rpo.NumOfcolumns = igrid.IGNORECOLUMNCOUNT
}

func (rpo *XLXSProcessorOptions) TotalPasses() int {
	return rpo.Passes
}

func (rpo *XLXSProcessorOptions) HeaderRowIndex() int {
	return rpo.HeaderRowindex
}

func (rpo *XLXSProcessorOptions) NumOfColumns() int {
	return rpo.NumOfcolumns
}

// String the readable version of the options
func (rpo *XLXSProcessorOptions) String() string {
	return fmt.Sprintf("Total Row Passes: %d, Header Index: %d, Number of Columns Index: %d, Sheets: %s", rpo.TotalPasses(), rpo.HeaderRowIndex(), rpo.NumOfColumns(), strings.Join(rpo.Sheets, ","))
}
