package main

import (
	"github.com/tdrip/griddata/pkg/csv"
	gd "github.com/tdrip/griddata/pkg/data"
)

func main() {
	// specify the file and the action
	gdp := csv.NewRowParserDefaultAction("./header.csv", gd.NewRowAction("PrintAction", gd.PrintCellAction))
	defer gdp.Close()

	err := gdp.Execute()

	if err != nil {
		panic(err)
	}
}
