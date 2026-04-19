package csv

import (
	"errors"
	"testing"

	gd "github.com/tdrip/griddata/pkg"
	igrid "github.com/tdrip/griddata/pkg/interfaces"
)

func TestCSV(t *testing.T) {
	gdp := CreateFileParser("../../testdata/noheader.csv")
	defer gdp.Close()

	err := gdp.Execute()

	if err != nil {
		t.Errorf("%s  ", err.Error())
	}

}

func TestCSVActions(t *testing.T) {
	csvtest := gd.CreateRowAction("PrintAction", gd.PrintCellAction)

	gdp := CreateFileParserWithAction("../../testdata/noheader.csv", &csvtest)
	defer gdp.Close()

	err := gdp.Execute()

	if err != nil {
		t.Errorf("%s  ", err.Error())
	}
}

func TestCSV3Passes(t *testing.T) {
	csvtest := gd.CreateRowAction("PrintAction", gd.PrintCellAction)

	gdp := CreateFileParserWithAction("../../testdata/noheader.csv", &csvtest)
	rowprocessors := gdp.GetProcessors()

	opts := rowprocessors[0].GetOptions()
	rpo := opts.(*gd.RowProcessorOptions)
	rpo.TotalPasses = 3
	rowprocessors[0].SetOptions(rpo)

	gdp.SetProcessors(rowprocessors)
	defer gdp.Close()

	err := gdp.Execute()

	if err != nil {
		t.Errorf("%s  ", err.Error())
	}

}

func FailAction(cell igrid.ICell) error {
	return errors.New("I should fail")
}
