package main

import (
	gd "github.com/tdrip/griddata/pkg"
	"github.com/tdrip/griddata/pkg/xlsx"
)

func main() {

	// this is a built in printer
	xlsxtest := gd.CreateRowAction("PrintAction", gd.PrintCellAction)

	// specify the file and the action
	gdp := xlsx.CreateFileParserWithAction("./header.xlsx", &xlsxtest)
	defer gdp.Close()

	err := gdp.Execute()

	if err != nil {
		panic(err)
	}
}
