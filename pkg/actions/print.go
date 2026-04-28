package actions

import (
	"fmt"

	igrid "github.com/tdrip/griddata/pkg/grid/interfaces"
)

func PrintCell(cell igrid.ICell) error {
	fmt.Println(cell.String())
	return nil
}

func PrintHeadedCell(header igrid.ICell, cell igrid.ICell) error {
	if header == nil {
		fmt.Println(cell.String())
	} else {
		fmt.Printf("Cell (%s) at [%s] has %v", header.GetData(), cell.GetLocation(), cell.GetData())
	}
	return nil
}
