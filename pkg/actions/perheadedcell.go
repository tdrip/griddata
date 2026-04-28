package actions

import (
	"errors"
	"fmt"

	uuid "github.com/google/uuid"
	iaction "github.com/tdrip/griddata/pkg/actions/interfaces"
	grid "github.com/tdrip/griddata/pkg/grid"
	igrid "github.com/tdrip/griddata/pkg/grid/interfaces"
)

type PerHeadedCellFunc func(igrid.ICell, igrid.ICell) error

// Headed Row Action An action that occurs on a Row
type PerHeadedCell struct {
	iaction.Action
	ID         string
	CellAction PerHeadedCellFunc
	Header     igrid.IRow
}

// creates a uuid for each action
func NewSimplePerHeadedCell(act PerHeadedCellFunc) PerHeadedCell {
	return PerHeadedCell{
		ID:         uuid.NewString(),
		CellAction: act,
	}
}

func NewPerHeadedCell(id string, act PerHeadedCellFunc) PerHeadedCell {
	return PerHeadedCell{
		ID:         id,
		CellAction: act,
	}
}

// Set Header
func (hra *PerHeadedCell) SetHeader(header igrid.IRow) {
	hra.Header = header
}

func (hra *PerHeadedCell) HasHeader() bool {
	return hra.Header != nil
}

// Get Id for this action
func (hra *PerHeadedCell) GetId() string {
	return hra.ID
}

func (hra *PerHeadedCell) Perform(data any) error {

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
