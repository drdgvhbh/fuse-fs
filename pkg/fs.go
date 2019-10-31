package pkg

import (
	"bazil.org/fuse/fs"
)

// FS -- Filesystem
type FS struct {
	root *Dir
}

// NewFs -- Create a new filesystem
func NewFs() *FS {
	iNodeGenerator := NewINodeGenerator(0)
	return &FS{
		root: NewDir(iNodeGenerator.next, iNodeGenerator),
	}
}

// Root - The root of the file system
func (f FS) Root() (fs.Node, error) {
	return f.root, nil
}
