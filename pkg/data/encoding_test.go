package data

import (
	"strings"
	"testing"
)

type TestRowHData struct {
	Column1         string   `row:"column1"`
	Column2         string   `row:"column2"`
	Column3         string   `row:"column3"`
	Column4         []string `row:"column4"`
	IgnoreThisField string
}

// column1,column2,column3,column4
func TestRowHDataDecode(t *testing.T) {

	trd := TestRowHData{
		Column1:         "NOTSET",
		Column2:         "NOTSET",
		Column3:         "NOTSET",
		Column4:         []string{},
		IgnoreThisField: "ignored",
	}

	rowdata, err := makeHeaderData()
	if err != nil {
		t.Errorf("makeHeaderData failed with %s", err.Error())
	}

	err = DecodeHeaderRowData(rowdata, &trd)
	if err != nil {
		t.Errorf("DecodeHeaderRowData failed with %s", err.Error())
	}

	if !strings.Contains(trd.Column1, "col1") {
		t.Errorf("column did not contain correct string (col1) has %s", trd.Column1)
	}

	if !strings.Contains(trd.Column2, "col2") {
		t.Errorf("column did not contain correct string (col2) has %s", trd.Column2)
	}

	if !strings.Contains(trd.Column3, "col3") {
		t.Errorf("column did not contain correct string (col3) has %s", trd.Column3)
	}

	if len(trd.Column4) == 0 {
		t.Errorf("column did not contain correct string array (col4) has %d value", len(trd.Column4))
	}

	for _, v := range trd.Column4 {
		if !strings.Contains(v, "col4") {
			t.Errorf("column did not contain correct string (col4) has %s", trd.Column4)
		}
	}

}

type TestIndexHRowData struct {
	Column1         string   `colindex:"0"`
	Column2         string   `colindex:"1"`
	Column3         string   `colindex:"2"`
	Column4         []string `colindex:"3"`
	IgnoreThisField string
}

func TestIndexHDataDecode(t *testing.T) {
	trd := TestIndexHRowData{
		Column1:         "NOTSET",
		Column2:         "NOTSET",
		Column3:         "NOTSET",
		Column4:         []string{},
		IgnoreThisField: "ignored",
	}

	rowdata, err := makeHeaderData()
	if err != nil {
		t.Errorf("makeHeaderData failed with %s", err.Error())
	}

	err = DecodeHeaderRowData(rowdata, &trd)
	if err != nil {
		t.Errorf("DecodeHeaderRowData failed with %s", err.Error())
	}

	if !strings.Contains(trd.Column1, "col1") {
		t.Errorf("column did not contain correct string (col1) has %s", trd.Column1)
	}

	if !strings.Contains(trd.Column2, "col2") {
		t.Errorf("column did not contain correct string (col2) has %s", trd.Column2)
	}

	if !strings.Contains(trd.Column3, "col3") {
		t.Errorf("column did not contain correct string (col3) has %s", trd.Column3)
	}

	if len(trd.Column4) == 0 {
		t.Errorf("column did not contain correct string array (col4) has %d value", len(trd.Column4))
	}

	for _, v := range trd.Column4 {
		if !strings.Contains(v, "col4") {
			t.Errorf("column did not contain correct string (col4) has %s", trd.Column4)
		}
	}

}

type TestIndexHNZRowData struct {
	Column1         string   `colindex:"1,nonzero"`
	Column2         string   `colindex:"2,nonzero"`
	Column3         string   `colindex:"3,nonzero"`
	Column4         []string `colindex:"4,nonzero"`
	IgnoreThisField string
}

func TestIndexHNZHeaderRowDataDecode(t *testing.T) {

	trd := TestIndexHNZRowData{
		Column1:         "NOTSET",
		Column2:         "NOTSET",
		Column3:         "NOTSET",
		Column4:         []string{},
		IgnoreThisField: "ignored",
	}

	rowdata, err := makeHeaderData()
	if err != nil {
		t.Errorf("makeHeaderData failed with %s", err.Error())
	}

	err = DecodeHeaderRowData(rowdata, &trd)
	if err != nil {
		t.Errorf("DecodeHeaderRowData failed with %s", err.Error())
	}

	if !strings.Contains(trd.Column1, "col1") {
		t.Errorf("column did not contain correct string (col1) has %s", trd.Column1)
	}

	if !strings.Contains(trd.Column2, "col2") {
		t.Errorf("column did not contain correct string (col2) has %s", trd.Column2)
	}

	if !strings.Contains(trd.Column3, "col3") {
		t.Errorf("column did not contain correct string (col3) has %s", trd.Column3)
	}

	if len(trd.Column4) == 0 {
		t.Errorf("column did not contain correct string array (col4) has %d value", len(trd.Column4))
	}

	for _, v := range trd.Column4 {
		if !strings.Contains(v, "col4") {
			t.Errorf("column did not contain correct string (col4) has %s", trd.Column4)
		}
	}

}

func makeHeaderData() (*HeaderRowData, error) {
	header := FillRowStringData(0, 1, []string{"column1", "column2", "column3", "column4"})
	return FillHeaderRowStringData(1, 1, []string{"col1row1", "col2row1", "col3row1", "\"col4row1,col4row1\""}, header)
}

func TestIndexHNZRowDataDecode(t *testing.T) {

	trd := TestIndexHNZRowData{
		Column1:         "NOTSET",
		Column2:         "NOTSET",
		Column3:         "NOTSET",
		Column4:         []string{},
		IgnoreThisField: "ignored",
	}

	rowdata := FillRowStringData(0, 1, []string{"col1row1", "col2row1", "col3row1", "\"col4row1,col4row1\""})

	err := DecodeRowData(rowdata, &trd)
	if err != nil {
		t.Errorf("DecodeHeaderRowData failed with %s", err.Error())
	}

	if !strings.Contains(trd.Column1, "col1") {
		t.Errorf("column did not contain correct string (col1) has %s", trd.Column1)
	}

	if !strings.Contains(trd.Column2, "col2") {
		t.Errorf("column did not contain correct string (col2) has %s", trd.Column2)
	}

	if !strings.Contains(trd.Column3, "col3") {
		t.Errorf("column did not contain correct string (col3) has %s", trd.Column3)
	}

	if len(trd.Column4) == 0 {
		t.Errorf("column did not contain correct string array (col4) has %d value", len(trd.Column4))
	}

	for _, v := range trd.Column4 {
		if !strings.Contains(v, "col4") {
			t.Errorf("column did not contain correct string (col4) has %s", trd.Column4)
		}
	}

}

func TestIndexDataDecode(t *testing.T) {
	trd := TestIndexHRowData{
		Column1:         "NOTSET",
		Column2:         "NOTSET",
		Column3:         "NOTSET",
		Column4:         []string{},
		IgnoreThisField: "ignored",
	}

	rowdata := FillRowStringData(0, 1, []string{"col1row1", "col2row1", "col3row1", "\"col4row1,col4row1\""})

	err := DecodeRowData(rowdata, &trd)
	if err != nil {
		t.Errorf("DecodeHeaderRowData failed with %s", err.Error())
	}

	if !strings.Contains(trd.Column1, "col1") {
		t.Errorf("column did not contain correct string (col1) has %s", trd.Column1)
	}

	if !strings.Contains(trd.Column2, "col2") {
		t.Errorf("column did not contain correct string (col2) has %s", trd.Column2)
	}

	if !strings.Contains(trd.Column3, "col3") {
		t.Errorf("column did not contain correct string (col3) has %s", trd.Column3)
	}

	if len(trd.Column4) == 0 {
		t.Errorf("column did not contain correct string array (col4) has %d value", len(trd.Column4))
	}

	for _, v := range trd.Column4 {
		if !strings.Contains(v, "col4") {
			t.Errorf("column did not contain correct string (col4) has %s", trd.Column4)
		}
	}

}
