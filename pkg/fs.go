package pkg

import (
	"bazil.org/fuse/fs"
)

type FS struct {
	root Dir
}

func NewFs() *FS {
	return &FS{
		root: Dir{},
	}
}

func (f FS) Root() (fs.Node, error) {
	return f.root, nil
}
