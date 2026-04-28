package actions

import (
	"errors"
	"fmt"

	iaction "github.com/tdrip/griddata/pkg/actions/interfaces"
	grid "github.com/tdrip/griddata/pkg/grid"
	igrid "github.com/tdrip/griddata/pkg/grid/interfaces"
)

type HeadedCellActionFunc func(igrid.ICell, igrid.ICell) error

// Headed Row Action An action that occurs on a Row
type HeadedCellAction struct {
	iaction.Action
	ID         string
	CellAction HeadedCellActionFunc
	Header     igrid.IRow
}

func NewHeadedCellAction(id string, act HeadedCellActionFunc) HeadedCellAction {
	return HeadedCellAction{
		ID:         id,
		CellAction: act,
	}
}

// Set Header
func (hra *HeadedCellAction) SetHeader(header igrid.IRow) {
	hra.Header = header
}

func (hra *HeadedCellAction) HasHeader() bool {
	return hra.Header != nil
}

// Get Id for this action
func (hra *HeadedCellAction) GetId() string {
	return hra.ID
}

func (hra *HeadedCellAction) Perform(data any) error {

	if hra.CellAction == nil {
		return fmt.Errorf("No action set for %s", hra.ID)
	}

	if data == nil {
		return fmt.Errorf("Row data was nil %s", hra.ID)
	}

	if hra.Header == nil {
		return fmt.Errorf("No Header data set for %s", hra.ID)
	}

	// We expect datarow to be correct type
	datarow, ok := data.(igrid.IRow)
	if !ok {
		return errors.New("data type was not Row Data")
	}
	for _, cell := range datarow.GetCells() {
		for _, hcell := range hra.Header.GetCells() {
			// same column
			if grid.MatchesY(hcell.GetLocation(), cell.GetLocation()) {
				err := hra.CellAction(hcell, cell)
				if err != nil {
					return err
				}
				break
			}
		}
	}
	return nil
}
