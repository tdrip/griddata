package main

import (
	gd "github.com/tdrip/griddata/pkg"
	"github.com/tdrip/griddata/pkg/xlsx"
)

func main() {
	// specify the file and the action
	gdp := xlsx.CreateRowParserWithAction("./header.xlsx", gd.CreateRowAction("PrintAction", gd.PrintCellAction))
	defer gdp.Close()

	err := gdp.Execute()

	if err != nil {
		panic(err)
	}
}
