package grid

import (
	igrid "github.com/tdrip/griddata/pkg/interfaces"
)

// HeadedRowData a row with a header
type HeadedRowData struct {
	igrid.IHeadedRow

	Header igrid.IHeader
}

// GetHeader Returns the header for the row
func (hrd *HeadedRowData) GetHeader() igrid.IHeader {
	return hrd.Header
}

// SetHeader Sets the header for the row
func (hrd *HeadedRowData) SetHeader(header igrid.IHeader) {
	hrd.Header = header
}
