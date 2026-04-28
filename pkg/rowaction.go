package grid

import (
	"errors"
	"fmt"

	uuid "github.com/google/uuid"
	idata "github.com/tdrip/griddata/pkg/data/interfaces"
)

// RowAction An action that occurs on a Row
type RowAction struct {
	idata.IDataAction
	ID     string
	Action CellAction
}

// creates a uuid for each action
func NewSimpleRowAction(act CellAction) RowAction {
	return RowAction{
		ID:     uuid.NewString(),
		Action: act,
	}
}

// Applied to every cell of the of the row
func NewRowAction(id string, act CellAction) RowAction {
	return RowAction{
		ID:     id,
		Action: act,
	}
}

// Get Id for this action
func (ra *RowAction) GetId() string {
	return ra.ID
}

func (ra *RowAction) Perform(data any) error {

	if ra.Action == nil {
		return fmt.Errorf("No action set for %s", ra.ID)
	}

	if data == nil {
		return fmt.Errorf("Row data was nil %s", ra.ID)
	}

	// We expect datarow to be correct type
	datarow, ok := data.(*RowData)
	if !ok {
		return errors.New("data type was not Row Data")
	}
	if datarow != nil {
		for _, cell := range datarow.GetCells() {
			err := ra.Action(cell)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
