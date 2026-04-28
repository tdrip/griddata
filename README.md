# griddata


## Benefits of this Library

### Row Actions
Row actions can be overridden and can be shared by different filetypes.
Thus a rowaction for a CSV can be written for a XSLX file.

### Encoding

```
type TestRowHData struct {
	Column1 string   `row:"column1"`
	Column2 string   `row:"column2"`
	Column3 string   `row:"column3"`
	Column4 []string `row:"column4"`
}
```

An example where a struct can be decoded per row based on the column headers

```
func testheaderdecode(rowdata *gd.HeaderRowData) error {
	trd := TestRowHData{}
	err := gd.DecodeHeaderRowData(rowdata, &trd)
	if err != nil {
		return err
	}

    return nil
}
```

## Examples

This example reads a csv with a header

```
package main

import (
	act "github.com/tdrip/griddata/pkg/actions"
	"github.com/tdrip/griddata/pkg/csv"
)

func main() {
	printcell := act.NewPerCellAction("print", act.PrintCellAction)

	// this has default options like seperator being a comma (,)
	gdp := csv.NewRowParserDefaultAction("./noheader.csv", &printcell)
	defer gdp.Close()

	err := gdp.Execute()

	if err != nil {
		panic(err)
	}
}
```

This is the example print out
```
Cell at [X:0,Y:0] has column1
Cell at [X:0,Y:1] has column2
Cell at [X:0,Y:2] has column3
Cell at [X:0,Y:3] has column4
Cell at [X:1,Y:0] has col1row1
Cell at [X:1,Y:1] has col2row1
Cell at [X:1,Y:2] has col3row1
Cell at [X:1,Y:3] has col4row1,col4row1
Cell at [X:2,Y:0] has col1row2
Cell at [X:2,Y:1] has col2row2
Cell at [X:2,Y:2] has col3row2
Cell at [X:2,Y:3] has col4row2,col4row2
Cell at [X:3,Y:0] has col1row3
Cell at [X:3,Y:1] has col2row3
Cell at [X:3,Y:2] has col3row3
Cell at [X:3,Y:3] has col4row3,col4row3
```
