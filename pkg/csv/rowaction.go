package csv

import (
	"fmt"

	gd "github.com/tdrip/griddata/pkg"
	igrid "github.com/tdrip/griddata/pkg/interfaces"
)

type CSVAction func(igrid.ICell) error

// CSVRowAction An action that occurs on a CSV Row
type CSVRowAction struct {
	igrid.IDataAction
	ID     string
	Action CSVAction
}

func PrintAction(cell igrid.ICell) error {
	fmt.Println(cell.String())
	return nil
}

func CreateCSVAction(id string, act CSVAction) CSVRowAction {
	return CSVRowAction{
		ID:     id,
		Action: act,
	}
}

// Get Id for this action
func (ra *CSVRowAction) GetId() string {
	return ra.ID
}

func (ra *CSVRowAction) PeformAction(data any) error {
	if ra.Action == nil {
		return fmt.Errorf("No action set for %s", ra.ID)
	}
	// We expect datarow to be correct tupe
	datarow := data.(*gd.RowData)
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
