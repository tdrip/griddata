package actions

import (
	"fmt"

	igrid "github.com/tdrip/griddata/pkg/grid/interfaces"
)

func PrintCell(cell igrid.Cell) error {
	fmt.Println(cell.String())
	return nil
}

func PrintHeadedCell(header igrid.Cell, cell igrid.Cell) error {
	if header == nil {
		fmt.Println(cell.String())
	} else {
		fmt.Printf("Cell (%s) at [%s] has %v", header.GetData(), cell.GetLocation(), cell.GetData())
	}
	return nil
}
