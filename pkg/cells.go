package grid

import (
	"fmt"

	igrid "github.com/tdrip/griddata/pkg/interfaces"
)

//Cell represents a cell of data
type Cell struct {
	igrid.ICell

	// Location of the Cell
	Location igrid.IPoint

	// data in the cell
	Data interface{}

	// any other cell data that should be related
	RelatedData map[igrid.IIndex]interface{}
}

//CreateStringCell creates a cell with string data
func CreateStringCell(location igrid.IPoint, data string) *Cell {
	gdc := Cell{}
	gdc.Location = location
	gdc.Data = data
	return &gdc
}

//GetLocation Gets the position
func (cell *Cell) GetLocation() igrid.IPoint {
	return cell.Location
}

//SetLocation Sets the position
func (cell *Cell) SetLocation(point igrid.IPoint) {
	cell.Location = point
}

//GetData Gets the data for the cell
func (cell *Cell) GetData() interface{} {
	return cell.Data
}

//SetData Sets the Data for the cell
func (cell *Cell) SetData(data interface{}) {
	//cell.Data = data.(string)
	cell.Data = data
}

//String Makes a respresentation of cell at point X:,Y: has data
func (cell Cell) String() string {
	return fmt.Sprintf("Cell at[%s] has %v", cell.GetLocation(), cell.GetData())
}
