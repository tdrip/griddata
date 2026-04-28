package main

import (
	act "github.com/tdrip/griddata/pkg/actions"
	"github.com/tdrip/griddata/pkg/csv"
)

func main() {
	// specify the file and the action
	gdp := csv.NewRowParserDefaultAction("./header.csv", act.NewPerCellAction("PrintAction", act.PrintCell))
	defer gdp.Close()

	err := gdp.Execute()

	if err != nil {
		panic(err)
	}
}
