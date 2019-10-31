package pkg

import (
	"context"
	"os"
	"time"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

const (
	// DirectoryPermission - Default Directory Permission on most Unix systems
	DirectoryPermission os.FileMode = 0755

	// DefaultSize - The default size in bytes
	DefaultSize uint64 = 4096
)

type INodeGenerator interface {
	Next() uint64
}

// Dir implements both Node and Handle for the root directory.
type Dir struct {
	iNode          uint64
	iNodeGenerator INodeGenerator
	size           uint64
	aTime          time.Time
	mTime          time.Time
	cTime          time.Time
}

func NewDir(iNode uint64, iNodeGenerator INodeGenerator) *Dir {
	return &Dir{
		iNode:          iNode,
		iNodeGenerator: iNodeGenerator,
		size:           DefaultSize,
		aTime:          time.Now(),
		mTime:          time.Now(),
		cTime:          time.Now(),
	}
}

func (d Dir) Attr(ctx context.Context, a *fuse.Attr) error {
	a.Inode = d.iNode
	a.Size = d.size
	a.Mode = os.ModeDir | DirectoryPermission
	a.Atime = d.aTime
	a.Mtime = d.mTime
	a.Ctime = d.cTime

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
