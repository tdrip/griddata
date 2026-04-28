package igrid

const (
	DEFAULTNUMOFPASSES = 1
	NOHEADERROWINDEX   = -1
	IGNORECOLUMNCOUNT  = -1
)

// IDataProcessorOptions Represents the options for parsing a data row
type IDataProcessorOptions interface {
	String() string

	// Pass over the data this many times
	// default should be DEFAULTNUMOFPASSES
	TotalPasses() int

	// default to no header
	// to ignore set to NOHEADERROWINDEX
	HeaderRowIndex() int

	// expected number of columns
	// if a data source should have 3 columns set to 3
	// to ignore set to IGNORECOLUMNCOUNT
	NumOfColumns() int
}
