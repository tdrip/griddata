package main

import (
	gd "github.com/tdrip/griddata/pkg"
	"github.com/tdrip/griddata/pkg/csv"
)

func main() {
	// specify the file and the action
	gdp := csv.NewRowParserWithDefaultHeaderAction("./header.csv", gd.NewRowAction("PrintAction", gd.PrintCellAction))
	defer gdp.Close()

	err := gdp.Execute()

	if err != nil {
		panic(err)
	}
}
