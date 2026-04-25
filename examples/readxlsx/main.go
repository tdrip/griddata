package main

import (
	gd "github.com/tdrip/griddata/pkg"
	"github.com/tdrip/griddata/pkg/xlsx"
)

func main() {
	xlsxtest := gd.CreateRowAction("PrintAction", gd.PrintCellAction)

	gdp := xlsx.CreateFileParserWithAction("./header.xlsx", &xlsxtest)
	defer gdp.Close()

	err := gdp.Execute()

	if err != nil {
		panic(err)
	}
}
