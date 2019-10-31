package pkg

import (
	"bazil.org/fuse/fs"
)

type FS struct{}

func (FS) Root() (fs.Node, error) {
	return Dir{}, nil
}
