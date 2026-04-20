package grid

import (
	"fmt"

	igrid "github.com/tdrip/griddata/pkg/interfaces"
)

type HeadedCellActionFunc func(igrid.ICell, igrid.ICell) error

// Headed Row Action An action that occurs on a Row
type HeadedCellAction struct {
	igrid.IDataAction
	ID     string
	Action HeadedCellActionFunc
	Header *RowData
}

func CreateHeadedCellAction(id string, act HeadedCellActionFunc) HeadedCellAction {
	return HeadedCellAction{
		ID:     id,
		Action: act,
	}
}

// Set Header
func (hra *HeadedCellAction) SetHeader(header *RowData) {
	hra.Header = header
}

func (hra *HeadedCellAction) HasHeader() bool {
	return hra.Header == nil
}

// Get Id for this action
func (hra *HeadedCellAction) GetId() string {
	return hra.ID
}

func (hra *HeadedCellAction) Peform(data any) error {

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
	if datarow != nil {
		for _, cell := range datarow.GetCells() {
			for _, hcell := range hra.Header.GetCells() {
				// same column
				if hcell.GetLocation().GetY() == cell.GetLocation().GetY() {
					err := hra.Action(hcell, cell)
					if err != nil {
						return err
					}
					break
				}
			}
		}
	}

	return nil
}
