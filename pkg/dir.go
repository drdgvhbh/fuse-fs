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

type Dir struct {
	iNode          uint64
	iNodeGenerator INodeGenerator
	mode           os.FileMode
	size           uint64
	aTime          time.Time
	mTime          time.Time
	cTime          time.Time
	parent         fs.Node
}

var (
	_ = fs.Node(&Dir{})
	_ = fs.NodeStringLookuper(&Dir{})
	_ = fs.HandleReadDirAller(&Dir{})
	_ = fs.NodeSetattrer(&Dir{})
)

func NewDir(iNode uint64, iNodeGenerator INodeGenerator) *Dir {
	dir := NewDirWithParent(iNode, iNodeGenerator, nil)
	dir.parent = dir
	return dir
}

func NewDirWithParent(iNode uint64, iNodeGenerator INodeGenerator, parent fs.Node) *Dir {
	return &Dir{
		iNode:          iNode,
		iNodeGenerator: iNodeGenerator,
		mode:           os.ModeDir | DirectoryPermission,
		size:           DefaultSize,
		aTime:          time.Now(),
		mTime:          time.Now(),
		cTime:          time.Now(),
		parent:         parent,
	}
}

func (d Dir) Attr(ctx context.Context, a *fuse.Attr) error {
	a.Inode = d.iNode
	a.Size = d.size
	a.Mode = d.mode
	a.Atime = d.aTime
	a.Mtime = d.mTime
	a.Ctime = d.cTime

	return nil
}

func (d *Dir) Setattr(
	ctx context.Context,
	req *fuse.SetattrRequest,
	resp *fuse.SetattrResponse,
) error {
	if req.Valid.Mode() {
		d.mode = req.Mode
	}
	if req.Valid.Size() {
		d.size = req.Size
	}
	if req.Valid.Atime() {
		d.aTime = req.Atime
	}
	if req.Valid.Mtime() {
		d.mTime = req.Mtime
	}

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
