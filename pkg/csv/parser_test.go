package csv

import (
	"os"
	"testing"

	logr "github.com/sirupsen/logrus"
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
