package csv

import (
	"testing"

	sl "github.com/tdrip/logger/pkg"
)

func TestCSV(t *testing.T) {

	// Open a log
	slog := sl.NewApplicationLogger()

	// lets open a flie log using the session
	slog.OpenAllChannels()

	//defer the close till the shell has closed
	defer slog.CloseAllChannels()

	gdp := CreateFileParser(slog, "../../testdata/noheader.csv")

	err := gdp.Execute()

	if err != nil {
		t.Errorf("%s  ", err.Error())
	}

}
