package igrid

//IColumnParsingOptions options for column parsring
type IColumnParsingOptions interface {
	String() string
}

//IColumnParser This interface parses a row in the grid
type IColumnParser interface {
	// Parse the column
	Parse(parent IParser, data IDataSource) error
}
