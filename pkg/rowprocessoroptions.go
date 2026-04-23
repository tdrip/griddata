package grid

import (
	"fmt"

	igrid "github.com/tdrip/griddata/pkg/interfaces"
)

// RowParsingOptions number of passes etc
type RowProcessorOptions struct {

	// inherit from the row IRowProcessingOptions
	igrid.IDataProcessorOptions

	TotalPasses int `json:"total_passes"`

	HeaderRowIndex int `json:"header_row_index"`

	NumOfColumns int `json:"number_columns"`
}

// Defaults
func (rpo *RowProcessorOptions) Defaults() {
	//Only pass over the row once
	rpo.TotalPasses = 1

	// default to no header
	rpo.HeaderRowIndex = -1

	rpo.NumOfColumns = -1
}

// String the readable version of the options
func (rpo *RowProcessorOptions) String() string {
	return fmt.Sprintf("Total Row Passes: %d,Header Index: %d,Number of Columns Index: %d", rpo.TotalPasses, rpo.HeaderRowIndex, rpo.NumOfColumns)
}
