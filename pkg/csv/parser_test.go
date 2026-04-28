package csv

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	gd "github.com/tdrip/griddata/pkg"
	igrid "github.com/tdrip/griddata/pkg/interfaces"
)

func TestCSV(t *testing.T) {
	gdp := CreateRowParser("../../testdata/noheader.csv")
	defer gdp.Close()

	err := gdp.Execute()

	if err != nil {
		t.Errorf("%s  ", err.Error())
	}

}

func TestCSVActions(t *testing.T) {
	gdp := CreateRowParserWithAction("../../testdata/noheader.csv",
		CreateCSVProcessorOptions(','),
		gd.CreateRowAction("PrintAction", gd.PrintCellAction))
	defer gdp.Close()

	err := gdp.Execute()

	if err != nil {
		t.Errorf("%s  ", err.Error())
	}
}

func TestCSV3Passes(t *testing.T) {
	opts := CreateCSVProcessorOptions(',')
	opts.Passes = 3
	gdp := CreateRowParserWithAction("../../testdata/noheader.csv", opts, gd.CreateRowAction("PrintAction", gd.PrintCellAction))
	rowprocessors := gdp.GetProcessors()

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

func TestCSVHeaderActions(t *testing.T) {
	opts := CreateCSVProcessorOptions(',')
	opts.HeaderRowindex = 0

	gdp := CreateRowParserWithHeaderAction("../../testdata/header.csv", opts, gd.CreateHeadedRowAction("testheader", testheader))
	defer gdp.Close()

	err := gdp.Execute()

	if err != nil {
		t.Errorf("%s  ", err.Error())
	}
}

// column1,column2,column3,column4
func testheader(rowdata *gd.HeaderRowData) error {
	s, err := rowdata.GetValString("column1")
	if err != nil {
		return err
	}

	if !strings.Contains(s, "col1") {
		return errors.New("wrong column should be 1")
	}

	s, err = rowdata.GetValString("column2")
	if err != nil {
		return err
	}

	if !strings.Contains(s, "col2") {
		return errors.New("wrong column should be 2")
	}

	s, err = rowdata.GetValString("column3")
	if err != nil {
		return err
	}

	if !strings.Contains(s, "col3") {
		return errors.New("wrong column should be 3")
	}

	ss, err := rowdata.GetValStringArray("column4", ",")
	if err != nil {
		return err
	}

	for _, v := range ss {
		if !strings.Contains(v, "col4") {
			return errors.New("wrong column should be 4")
		}
	}
	//return errors.New("I failed")
	return nil
}

func TestCSVHeaderActionDecode(t *testing.T) {
	opts := CreateCSVProcessorOptions(',')
	opts.HeaderRowindex = 0
	gdp := CreateRowParserWithHeaderAction("../../testdata/header.csv", opts, gd.CreateHeadedRowAction("testheaderdecode", testheaderdecode))
	defer gdp.Close()

	err := gdp.Execute()

	if err != nil {
		t.Errorf("%s  ", err.Error())
	}
}

type TestRowHData struct {
	Column1 string   `row:"column1"`
	Column2 string   `row:"column2"`
	Column3 string   `row:"column3"`
	Column4 []string `row:"column4"`
}

// column1,column2,column3,column4
func testheaderdecode(rowdata *gd.HeaderRowData) error {

	trd := TestRowHData{
		Column1: "NOTSET",
		Column2: "NOTSET",
		Column3: "NOTSET",
		Column4: []string{},
	}
	err := gd.DecodeHeaderRowData(rowdata, &trd)
	if err != nil {
		return err
	}

	if !strings.Contains(trd.Column1, "col1") {
		return fmt.Errorf("column did not contain correct string (col1) has %s", trd.Column1)
	}

	if !strings.Contains(trd.Column2, "col2") {
		return fmt.Errorf("column did not contain correct string (col2) has %s", trd.Column2)
	}

	if !strings.Contains(trd.Column3, "col3") {
		return fmt.Errorf("column did not contain correct string (col3) has %s", trd.Column3)
	}

	if len(trd.Column4) == 0 {
		return fmt.Errorf("column did not contain correct string array (col4) has %d value", len(trd.Column4))
	}

	for _, v := range trd.Column4 {
		if !strings.Contains(v, "col4") {
			return fmt.Errorf("column did not contain correct string (col4) has %s", trd.Column4)
		}
	}

	return nil
}

type TestIndexHRowData struct {
	Column1 string   `colindex:"0"`
	Column2 string   `colindex:"1"`
	Column3 string   `colindex:"2"`
	Column4 []string `colindex:"3"`
}

func TestCSVIndexHeaderActionDecode(t *testing.T) {
	opts := CreateCSVProcessorOptions(',')
	opts.HeaderRowindex = 0
	gdp := CreateRowParserWithHeaderAction("../../testdata/header.csv", opts, gd.CreateHeadedRowAction("testheaderindexdecode", testheaderindexdecode))
	defer gdp.Close()

	err := gdp.Execute()

	if err != nil {
		t.Errorf("%s  ", err.Error())
	}
}

// column1,column2,column3,column4
func testheaderindexdecode(rowdata *gd.HeaderRowData) error {

	trd := TestIndexHRowData{
		Column1: "NOTSET",
		Column2: "NOTSET",
		Column3: "NOTSET",
		Column4: []string{},
	}
	err := gd.DecodeHeaderRowData(rowdata, &trd)
	if err != nil {
		return err
	}

	if !strings.Contains(trd.Column1, "col1") {
		return fmt.Errorf("column did not contain correct string (col1) has %s", trd.Column1)
	}

	if !strings.Contains(trd.Column2, "col2") {
		return fmt.Errorf("column did not contain correct string (col2) has %s", trd.Column2)
	}

	if !strings.Contains(trd.Column3, "col3") {
		return fmt.Errorf("column did not contain correct string (col3) has %s", trd.Column3)
	}

	if len(trd.Column4) == 0 {
		return fmt.Errorf("column did not contain correct string array (col4) has %d value", len(trd.Column4))
	}

	for _, v := range trd.Column4 {
		if !strings.Contains(v, "col4") {
			return fmt.Errorf("column did not contain correct string (col4) has %s", trd.Column4)
		}
	}

	return nil
}

type TestIndexHNZRowData struct {
	Column1 string   `colindex:"1,nonzero"`
	Column2 string   `colindex:"2,nonzero"`
	Column3 string   `colindex:"3,nonzero"`
	Column4 []string `colindex:"4,nonzero"`
}

func TestCSVNZIndexHeaderActionDecode(t *testing.T) {
	opts := CreateCSVProcessorOptions(',')
	opts.HeaderRowindex = 0
	gdp := CreateRowParserWithHeaderAction("../../testdata/header.csv", opts, gd.CreateHeadedRowAction("testheadernzindexdecode", testheadernzindexdecode))
	defer gdp.Close()

	err := gdp.Execute()

	if err != nil {
		t.Errorf("%s  ", err.Error())
	}
}

// column1,column2,column3,column4
func testheadernzindexdecode(rowdata *gd.HeaderRowData) error {

	trd := TestIndexHNZRowData{
		Column1: "NOTSET",
		Column2: "NOTSET",
		Column3: "NOTSET",
		Column4: []string{},
	}
	err := gd.DecodeHeaderRowData(rowdata, &trd)
	if err != nil {
		return err
	}

	if !strings.Contains(trd.Column1, "col1") {
		return fmt.Errorf("column did not contain correct string (col1) has %s", trd.Column1)
	}

	if !strings.Contains(trd.Column2, "col2") {
		return fmt.Errorf("column did not contain correct string (col2) has %s", trd.Column2)
	}

	if !strings.Contains(trd.Column3, "col3") {
		return fmt.Errorf("column did not contain correct string (col3) has %s", trd.Column3)
	}

	if len(trd.Column4) == 0 {
		return fmt.Errorf("column did not contain correct string array (col4) has %d value", len(trd.Column4))
	}

	for _, v := range trd.Column4 {
		if !strings.Contains(v, "col4") {
			return fmt.Errorf("column did not contain correct string (col4) has %s", trd.Column4)
		}
	}

	return nil
}

func TestTSVNZIndexHeaderActionDecode(t *testing.T) {
	opts := CreateCSVProcessorOptions('	')
	opts.HeaderRowindex = 0
	gdp := CreateRowParserWithHeaderAction("../../testdata/header.tsv", opts, gd.CreateHeadedRowAction("testheadernzindexdecode", testheadernzindexdecode))
	defer gdp.Close()

	err := gdp.Execute()

	if err != nil {
		t.Errorf("%s  ", err.Error())
	}
}

func TestTSVIndexHeaderActionDecode(t *testing.T) {
	opts := CreateCSVProcessorOptions('	')
	opts.HeaderRowindex = 0
	gdp := CreateRowParserWithHeaderAction("../../testdata/header.tsv", opts, gd.CreateHeadedRowAction("testheaderindexdecode", testheaderindexdecode))
	defer gdp.Close()

	err := gdp.Execute()

	if err != nil {
		t.Errorf("%s  ", err.Error())
	}
}

func TestTSVHeaderActionDecode(t *testing.T) {
	opts := CreateCSVProcessorOptions('	')
	opts.HeaderRowindex = 0
	gdp := CreateRowParserWithHeaderAction("../../testdata/header.tsv", opts, gd.CreateHeadedRowAction("testheaderdecode", testheaderdecode))
	defer gdp.Close()

	err := gdp.Execute()

	if err != nil {
		t.Errorf("%s  ", err.Error())
	}
}

func TestTSVActions(t *testing.T) {
	gdp := CreateRowParserWithAction("../../testdata/noheader.tsv",
		CreateCSVProcessorOptions('	'),
		gd.CreateRowAction("PrintAction", gd.PrintCellAction))
	defer gdp.Close()

	err := gdp.Execute()

	if err != nil {
		t.Errorf("%s  ", err.Error())
	}
}

func TestTSV3Passes(t *testing.T) {
	opts := CreateCSVProcessorOptions('	')
	opts.Passes = 3
	gdp := CreateRowParserWithAction("../../testdata/noheader.tsv", opts, gd.CreateRowAction("PrintAction", gd.PrintCellAction))
	rowprocessors := gdp.GetProcessors()

	gdp.SetProcessors(rowprocessors)
	defer gdp.Close()

	err := gdp.Execute()

	if err != nil {
		t.Errorf("%s  ", err.Error())
	}

}
