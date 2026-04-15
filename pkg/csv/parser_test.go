package csv

import (
	"errors"
	"fmt"
	"os"
	"testing"

	logr "github.com/sirupsen/logrus"
	igrid "github.com/tdrip/griddata/pkg/interfaces"
)

func TestCSV(t *testing.T) {

	log := logr.New()

	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&logr.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(logr.TraceLevel)

	gdp := CreateFileParser(log, "../../testdata/noheader.csv")

	err := gdp.Execute()

	if err != nil {
		t.Errorf("%s  ", err.Error())
	}

}

func TestCSVActions(t *testing.T) {

	log := logr.New()

	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&logr.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(logr.TraceLevel)

	csvtest := CreateCSVAction("PrintAction", PrintAction)

	gdp := CreateFileParserWithAction(log, "../../testdata/noheader.csv", &csvtest)

	err := gdp.Execute()

	if err != nil {
		t.Errorf("%s  ", err.Error())
	}

}

func TestCSV3Passes(t *testing.T) {

	log := logr.New()

	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&logr.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(logr.TraceLevel)

	csvtest := CreateCSVAction("PrintAction", PrintAction)

	gdp := CreateFileParserWithAction(log, "../../testdata/noheader.csv", &csvtest)
	rowprocessors := gdp.GetProcessors()

	opts := rowprocessors[0].GetOptions()
	rpo := opts.(*RowProcessorOptions)
	rpo.TotalPasses = 3
	rowprocessors[0].SetOptions(rpo)

	gdp.SetProcessors(rowprocessors)

	err := gdp.Execute()

	if err != nil {
		t.Errorf("%s  ", err.Error())
	}

}

func FailAction(cell igrid.ICell) error {
	return errors.New("I should fail")
}

func PrintAction(cell igrid.ICell) error {
	fmt.Println(cell.String())
	return nil
}
