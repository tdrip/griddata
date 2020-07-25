package datasources

import (
	"encoding/csv"
	"os"

	igrid "github.com/tdrip/griddata/pkg/interfaces"
)

//CSVFile represents a CSV file
type CSVFile struct {
	igrid.IDataSource

	Filepath   string
	filestream *os.File
	reader     *csv.Reader
}

//CreateCSVFile Creates a CSV file
func CreateCSVFile(filepath string) *CSVFile {
	csvsource := CSVFile{Filepath: filepath}
	return &csvsource
}

//Validate checks if the file exists
func (csvf *CSVFile) Validate() error {

	_, err := os.Stat(csvf.Filepath)
	if os.IsNotExist(err) {
		return err
	}

	return nil
}

//Open opens the file stream and creates csv reader
func (csvf *CSVFile) Open() error {

	f, err := os.Open(csvf.Filepath)
	if err != nil {
		return err
	}

	csvf.filestream = f
	csvf.reader = csv.NewReader(f)

	return nil
}

// Close closes the stream
func (csvf *CSVFile) Close() error {

	if csvf.filestream != nil {
		return csvf.filestream.Close()
	}

	return nil
}
