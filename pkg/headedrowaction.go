package grid

import (
	"errors"
	"fmt"

	igrid "github.com/tdrip/griddata/pkg/interfaces"
)

type HeadedRowActionFunc func(*HeaderRowData, *RowData) error

// Headed Row Action An action that occurs on a Row
type HeadedRowAction struct {
	igrid.IDataAction
	ID     string
	Action HeadedRowActionFunc
	Header *HeaderRowData
}

func CreateHeadedRowAction(id string, act HeadedRowActionFunc) HeadedRowAction {
	return HeadedRowAction{
		ID:     id,
		Action: act,
	}
}

// Get Id for this action
func (hra *HeadedRowAction) GetId() string {
	return hra.ID
}

func (hra *HeadedRowAction) Perform(data any) error {

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
	datarow, ok := data.(*RowData)
	if !ok {
		return errors.New("data type was not Row Data")
	}

	return hra.Action(hra.Header, datarow)
}
