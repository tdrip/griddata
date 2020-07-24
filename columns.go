package grid

import (
	"fmt"

	igrid "github.com/tdrip/griddata/pkg/interfaces"
)

//GDColumMapping This structs has the mapping of the column indexs
type GDColumMapping struct {
	ColumnIndexes []igrid.IIndex `json:"columnindexes,omitempty"`
}

//GDColumnIndex This struct is the column index
// it can be based around the name or an integer index
// the name will be read by the header row
type GDColumnIndex struct {
	igrid.IIndex

	Index int `json:"index"`

	Name string `json:"name"`

	// do we need to associate other columns with this column?
	// for example: column 1 is the first name and column 2 is the second name they should be associated
	RelatedIndexes []igrid.IIndex `json:"relatedindexes,omitempty"`
}

func (gdci *GDColumnIndex) GetPosition() int {
	return gdci.Index
}

func (gdci *GDColumnIndex) SetPosition(position int) {
	gdci.Index = position
}

func (gdci *GDColumnIndex) GetDisplayName() string {
	return gdci.Name
}

func (gdci *GDColumnIndex) SetDisplayName(displayname string) {
	gdci.Name = displayname
}

func (gdci *GDColumnIndex) GetRelatedIndexes() []igrid.IIndex {
	return gdci.RelatedIndexes
}

func (gdci *GDColumnIndex) SetRelatedIndexes(columns []igrid.IIndex) {
	gdci.RelatedIndexes = columns
}

func CreateRColumnIndexbyIndex(Index int, rcolumns []igrid.IIndex) igrid.IIndex {
	return CreateColumnIndex(Index, "", rcolumns)
}

func CreateRColumnIndexbyName(Name string, rcolumns []igrid.IIndex) igrid.IIndex {
	return CreateColumnIndex(-1, Name, rcolumns)
}

func CreateColumnIndexbyIndex(Index int) igrid.IIndex {
	return CreateColumnIndex(Index, "", nil)
}

func CreateColumnIndexbyName(Name string) igrid.IIndex {
	return CreateColumnIndex(-1, Name, nil)
}

// Create Column Index
func CreateColumnIndex(index int, name string, rcolumns []igrid.IIndex) igrid.IIndex {
	ci := GDColumnIndex{Index: index, Name: name, RelatedIndexes: rcolumns}
	return &ci
}

//GDColumnIndex Print ColumnIndex
func (ci GDColumnIndex) String() string {

	if ci.Name != "" {
		if ci.Index < 0 {
			return fmt.Sprintf("[%s]", ci.Name)
		} else {
			return fmt.Sprintf("[%s(%d)]", ci.Name, ci.Index)
		}
	} else {
		return fmt.Sprintf("[%d]", ci.Index)
	}

}
