package igrid

// IDataProcessorOptions Represents the options for parsing a data row
type IDataProcessorOptions interface {
	Defaults()
	String() string
}
