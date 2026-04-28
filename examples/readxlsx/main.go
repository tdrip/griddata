package main

import (
	"time"

	act "github.com/tdrip/griddata/pkg/actions"
	igrid "github.com/tdrip/griddata/pkg/grid/interfaces"
	"github.com/tdrip/griddata/pkg/xlsx"
)

func main() {
	// specify the file and the action

	// normal print
	//	gdp := xlsx.NewRowParserWithAction("./header.xlsx", act.NewPerCellAction("PrintAction", act.PrintCellAction))

	//  slow print
	gdp := xlsx.NewRowParserDefaultAction("./header.xlsx", act.NewPerCellAction("SlowPrint", SlowPrint))
	defer gdp.Close()

	err := gdp.Execute()

	if err != nil {
		panic(err)
	}
}

func SlowPrint(cell igrid.ICell) error {
	time.Sleep(500 * time.Millisecond)
	act.PrintCellAction(cell)
	return nil
}
