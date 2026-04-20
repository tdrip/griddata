package grid

import (
	"fmt"

	igrid "github.com/tdrip/griddata/pkg/interfaces"
)

type HeadedRowActionFunc func(*RowData, *RowData) error

// Headed Row Action An action that occurs on a Row
type HeadedRowAction struct {
	igrid.IDataAction
	ID     string
	Action HeadedRowActionFunc
	Header *RowData
}

func CreateHeadedRowAction(id string, act HeadedRowActionFunc) HeadedRowAction {
	return HeadedRowAction{
		ID:     id,
		Action: act,
	}
}

// Set Header
func (hra *HeadedRowAction) SetHeader(header *RowData) {
	hra.Header = header
}

func (hra *HeadedRowAction) HasHeader() bool {
	return hra.Header == nil
}

// Get Id for this action
func (hra *HeadedRowAction) GetId() string {
	return hra.ID
}

func (hra *HeadedRowAction) Peform(data any) error {

	if hra.Action == nil {
		return fmt.Errorf("No action set for %s", hra.ID)
	}

	if data == nil {
		return fmt.Errorf("Row data was nil %s", hra.ID)
	}

	if hra.Header == nil {
		return fmt.Errorf("No Header data set for %s", hra.ID)
	}

	// We expect datarow to be correct type
	datarow := data.(*RowData)

	if hra.Action != nil {
		return hra.Action(hra.Header, datarow)
	}
	return nil
}
