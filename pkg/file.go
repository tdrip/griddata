package grid

import (
	"fmt"
	"os"
	"path/filepath"

	igrid "github.com/tdrip/griddata/pkg/interfaces"
)

// Grid File represents a sandard file
// basic open,close and validation
type GridFile struct {
	igrid.IDataSource

	Filepath   string
	Filestream *os.File
}

// New Grid File to represent a file datasource
func NewGridFile(fp string) *GridFile {
	source := &GridFile{Filepath: filepath.Clean(fp)}
	return source
}

// Validate checks if the file exists
func (gf *GridFile) Validate() error {

	info, err := os.Stat(gf.Filepath)
	if os.IsNotExist(err) {
		return err
	}

	// empty file
	if info.Size() == 0 {
		return fmt.Errorf("file has %d file size", info.Size())
	}

	// directory
	if info.IsDir() {
		return fmt.Errorf("%s is a directory not a file", gf.Filepath)
	}

	return nil
}

// Open opens the file stream and creates csv reader
func (gf *GridFile) Open() error {

	f, err := os.Open(gf.Filepath)
	if err != nil {
		return err
	}

	gf.Filestream = f

	return nil
}

// Close closes the stream
func (gf *GridFile) Close() error {

	if gf.Filestream != nil {
		return gf.Filestream.Close()
	}

	return nil
}
