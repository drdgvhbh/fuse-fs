package pkg

import (
	"context"
	"os"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

// DirectoryPermission - Default Directory Permission on most Unix systems
const DirectoryPermission = 0755

type INodeGenerator interface {
	Next() uint64
}

// Dir implements both Node and Handle for the root directory.
type Dir struct {
	iNode          uint64
	iNodeGenerator INodeGenerator
}

func NewDir(iNode uint64, iNodeGenerator INodeGenerator) *Dir {
	return &Dir{
		iNode,
		iNodeGenerator,
	}
}

func (d Dir) Attr(ctx context.Context, a *fuse.Attr) error {
	a.Inode = d.iNode
	a.Mode = os.ModeDir | DirectoryPermission
	return nil
}

func (Dir) Lookup(ctx context.Context, name string) (fs.Node, error) {
	if name == "hello" {
		return File{}, nil
	}
	return nil, fuse.ENOENT
}

var dirDirs = []fuse.Dirent{
	{Inode: 2, Name: "hello", Type: fuse.DT_File},
}

func (Dir) ReadDirAll(ctx context.Context) ([]fuse.Dirent, error) {
	return dirDirs, nil
}
