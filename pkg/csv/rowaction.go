package csv

import (
	"fmt"
	gd "github.com/tdrip/griddata/pkg"
	igrid "github.com/tdrip/griddata/pkg/interfaces"
)

//CSVRowAction An action that occurs on a CSV Row
type CSVRowAction struct {
	igrid.IDataAction
	ID string
}

//Get Id for this action
func (ra *CSVRowAction) GetId() string {
	return ra.ID
}

func (ra *CSVRowAction) PeformAction(data interface{}) {
	// We datarow to be correct tupe
	datarow := data.(*gd.RowData)
	if datarow != nil {
		for _, cell := range datarow.GetCells() {
			fmt.Printf("PeformAction - %v", cell)
		}
	}
}
