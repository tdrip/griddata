package xlsx

import (
	"fmt"
	"os"
	"path/filepath"

	igrid "github.com/tdrip/griddata/pkg/interfaces"
)

// XLSXFile represents a XLSX file
type XLSXFile struct {
	igrid.IDataSource

	Filepath   string
	Filestream *os.File
	Sheets     []string
}

// CreateXLSXFile Creates a XLSX file
func CreateXLSXFile(fp string) *XLSXFile {
	source := XLSXFile{
		Filepath: filepath.Clean(fp),
		Sheets:   []string{},
	}
	return &source
}

func CreateXLSXFileSheet(fp string, sheetname string) *XLSXFile {
	source := XLSXFile{
		Filepath: filepath.Clean(fp),
		Sheets:   []string{sheetname},
	}
	return &source
}

func CreateXLSXFileSheets(fp string, sheetnames []string) *XLSXFile {
	source := XLSXFile{
		Filepath: filepath.Clean(fp),
		Sheets:   sheetnames,
	}
	return &source
}

// Validate checks if the file exists
func (xlsxf *XLSXFile) Validate() error {

	info, err := os.Stat(xlsxf.Filepath)
	if os.IsNotExist(err) {
		return err
	}

	// empty file
	if info.Size() == 0 {
		return fmt.Errorf("file has %d file size", info.Size())
	}

	// directory
	if info.IsDir() {
		return fmt.Errorf("%s is a directory not a file", xlsxf.Filepath)
	}

	return nil
}

// Open opens the file stream and creates csv reader
func (xlsxf *XLSXFile) Open() error {

	f, err := os.Open(filepath.Clean(xlsxf.Filepath))
	if err != nil {
		return err
	}
	xlsxf.Filestream = f
	return nil
}

// Close closes the stream
func (xlsxf *XLSXFile) Close() error {

	if xlsxf.Filestream != nil {
		return xlsxf.Filestream.Close()
	}

	return nil
}
