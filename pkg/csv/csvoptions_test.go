package csv

import (
	"testing"

	act "github.com/tdrip/griddata/pkg/actions"
)

func TestCSVOptions(t *testing.T) {

	headerindex := 0
	numberofpasses := 22
	numberofcolumns := 45

	gdp := NewHeadedRowParserWithAction("../../testdata/header.csv",
		act.NewSimpleHeadedRow(testheadernzindexdecode),
		WithHeaderIndex(headerindex))

	if gdp.Processors[0].GetOptions().HeaderRowIndex() != headerindex {
		t.Errorf("%d header row index were set when it should be %d", gdp.Processors[0].GetOptions().HeaderRowIndex(), headerindex)
	}

	gdp = NewHeadedRowParserWithAction("../../testdata/header.csv",
		act.NewSimpleHeadedRow(testheadernzindexdecode),
		WithHeaderIndex(headerindex),
		WithRowPasses(numberofpasses))

	if gdp.Processors[0].GetOptions().RowPasses() != numberofpasses {
		t.Errorf("%d passes were set when it should be %d", gdp.Processors[0].GetOptions().RowPasses(), numberofpasses)
	}

	gdp = NewHeadedRowParserWithAction("../../testdata/header.csv",
		act.NewSimpleHeadedRow(testheadernzindexdecode),
		WithHeaderIndex(headerindex),
		WithRowPasses(numberofpasses),
		WithMinColumns(numberofcolumns))

	if gdp.Processors[0].GetOptions().MinColumns() != numberofcolumns {
		t.Errorf("%d number of columns were set when it should be %d", gdp.Processors[0].GetOptions().MinColumns(), numberofcolumns)
	}

	// real test
	gdp = NewHeadedRowParserWithAction("../../testdata/header.csv",
		act.NewSimpleHeadedRow(testheadernzindexdecode),
		WithHeaderIndex(0),
		WithRowPasses(1),
		WithMinColumns(4))

	defer gdp.Close()

	err := gdp.Execute()

	if err != nil {
		t.Errorf("%s  ", err.Error())
	}

}
