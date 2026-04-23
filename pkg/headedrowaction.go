package grid

import (
	"errors"
	"fmt"

	igrid "github.com/tdrip/griddata/pkg/interfaces"
)

type HeadedRowActionFunc func(*HeaderRowData) error

// Headed Row Action An action that occurs on a Row
type HeadedRowAction struct {
	igrid.IDataAction
	ID     string
	Action HeadedRowActionFunc
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

	// We expect datarow to be correct type
	datarow, ok := data.(*HeaderRowData)
	if !ok {
		return errors.New("data type was not Headed Row Data")
	}

	return hra.Action(datarow)
}
