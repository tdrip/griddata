package main

import (
	"time"

	gd "github.com/tdrip/griddata/pkg"
	igrid "github.com/tdrip/griddata/pkg/interfaces"
	"github.com/tdrip/griddata/pkg/xlsx"
)

func main() {
	// specify the file and the action

	// normal print
	//	gdp := xlsx.CreateRowParserWithAction("./header.xlsx", gd.CreateRowAction("PrintAction", gd.PrintCellAction))

	//  slow print
	gdp := xlsx.CreateRowParserWithAction("./header.xlsx", gd.CreateRowAction("SlowPrint", SlowPrint))
	defer gdp.Close()

	err := gdp.Execute()

	if err != nil {
		panic(err)
	}
}

func SlowPrint(cell igrid.ICell) error {
	time.Sleep(500 * time.Millisecond)
	gd.PrintCellAction(cell)
	return nil
}
