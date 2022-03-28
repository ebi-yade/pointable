package pointable

import (
	"fmt"
	"os"
	"path/filepath"
)

// Dist represents the package name and directory path of artifacts.
type Dist struct {
	name string
	dir  string
}

// NewDist returns a brand-new Dist
func NewDist(path string) *Dist {
	return &Dist{
		name: filepath.Base(path),
		dir:  filepath.Clean(path),
	}
}

// SetName update the name of the distributed package.
func (d *Dist) SetName(name string) *Dist {
	d.name = name
	return d
}

// Validate checks if the Dist has correct values.
func (d *Dist) Validate() (*Dist, error) {
	_, err := os.Stat(d.dir)
	if err == nil {
		return d, fmt.Errorf("a directory or file '%s' already exists", d.dir)
	}
	if os.IsExist(err) {
		return d, fmt.Errorf("failed to check a directory or file '%s' does not exist: %w", d.dir, err)
	}

	// TODO: check package name if it match the pattern of Go packages' name
	return d, nil
}

func (d *Dist) createFile(name string) (*os.File, error) {
	if err := os.MkdirAll(d.dir, os.ModePerm); err != nil {
		return nil, fmt.Errorf("failed to make the directory '%s': %w", d.dir, err)
	}

	filePath := filepath.Join(d.dir, fmt.Sprintf("%s.go", name))
	f, err := os.Create(filePath)
	if err != nil {
		return f, fmt.Errorf("failed to createFile the file '%s': %w", filePath, err)
	}

	return f, nil
}
