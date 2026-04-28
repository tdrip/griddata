package actions

import (
	"fmt"

	uuid "github.com/google/uuid"
	iaction "github.com/tdrip/griddata/pkg/actions/interfaces"
	igrid "github.com/tdrip/griddata/pkg/grid/interfaces"
)

type HeadedRowActionFunc func(igrid.IRow) error

// Headed Row Action An action that occurs on a Row
type HeadedRowAction struct {
	iaction.Action
	ID        string
	RowAction HeadedRowActionFunc
}

// creates a uuid for each action
func NewSimpleHeadedRowAction(act HeadedRowActionFunc) HeadedRowAction {
	return HeadedRowAction{
		ID:        uuid.NewString(),
		RowAction: act,
	}
}

func NewHeadedRowAction(id string, act HeadedRowActionFunc) HeadedRowAction {
	return HeadedRowAction{
		ID:        id,
		RowAction: act,
	}
}

// Get Id for this action
func (hra *HeadedRowAction) GetId() string {
	return hra.ID
}

func (hra *HeadedRowAction) Perform(data any) error {

	if hra.RowAction == nil {
		return fmt.Errorf("No row action set for %s", hra.ID)
	}

	if data == nil {
		return fmt.Errorf("Row data was nil %s", hra.ID)
	}

	// We expect data to be the interface for the action
	// not tied to the data layer
	datarow, ok := data.(igrid.IRow)
	if !ok {
		return fmt.Errorf("data type was not Headed Row Data - Raw Data:  %v", data)
	}

	return hra.RowAction(datarow)
}
