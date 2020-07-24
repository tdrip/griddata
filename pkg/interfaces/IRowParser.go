package igrid

type IRowParsingOptions interface {
}

//IRowParser This interface parses a row in the grid
type IRowParser interface {
	// Parse the row
	Parse(parent IParser, data IDataSource) error
}
