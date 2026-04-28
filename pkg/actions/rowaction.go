package actions

import (
	"errors"
	"fmt"

	uuid "github.com/google/uuid"
	iaction "github.com/tdrip/griddata/pkg/actions/interfaces"
	igrid "github.com/tdrip/griddata/pkg/grid/interfaces"
)

type CellAction func(igrid.ICell) error

// RowAction An action that occurs on a Row
type RowAction struct {
	iaction.Action
	ID         string
	CellAction CellAction
}

// creates a uuid for each action
func NewSimpleRowAction(act CellAction) RowAction {
	return RowAction{
		ID:         uuid.NewString(),
		CellAction: act,
	}
}

// Applied to every cell of the of the row
func NewRowAction(id string, act CellAction) RowAction {
	return RowAction{
		ID:         id,
		CellAction: act,
	}
}

// Get Id for this action
func (ra *RowAction) GetId() string {
	return ra.ID
}

func (ra *RowAction) Perform(data any) error {

	if ra.CellAction == nil {
		return fmt.Errorf("No action set for %s", ra.ID)
	}

	if data == nil {
		return fmt.Errorf("Row data was nil %s", ra.ID)
	}

	// We expect data to be the interface for the action
	// not tied to the data layer
	datarow, ok := data.(igrid.IRow)
	if !ok {
		return errors.New("data type was not Row Data")
	}

	for _, cell := range datarow.GetCells() {
		err := ra.CellAction(cell)
		if err != nil {
			return err
		}
	}

	return nil
}
