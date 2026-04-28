package data

import (
	"fmt"

	uuid "github.com/google/uuid"
	idata "github.com/tdrip/griddata/pkg/data/interfaces"
)

type HeadedRowActionFunc func(*HeaderRowData) error

// Headed Row Action An action that occurs on a Row
type HeadedRowAction struct {
	idata.Action
	ID         string
	CellAction HeadedRowActionFunc
}

// creates a uuid for each action
func NewSimpleHeadedRowAction(act HeadedRowActionFunc) HeadedRowAction {
	return HeadedRowAction{
		ID:         uuid.NewString(),
		CellAction: act,
	}
}

func NewHeadedRowAction(id string, act HeadedRowActionFunc) HeadedRowAction {
	return HeadedRowAction{
		ID:         id,
		CellAction: act,
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
		return fmt.Errorf("data type was not Headed Row Data - Raw Data:  %v", data)
	}

	return hra.CellAction(datarow)
}
