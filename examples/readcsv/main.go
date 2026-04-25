package main

import (
	gd "github.com/tdrip/griddata/pkg"
	"github.com/tdrip/griddata/pkg/csv"
)

func main() {
	xlsxtest := gd.CreateRowAction("PrintAction", gd.PrintCellAction)

	gdp := csv.CreateFileParserWithAction("./header.csv", &xlsxtest)
	defer gdp.Close()

	err := gdp.Execute()

	if err != nil {
		panic(err)
	}
}
