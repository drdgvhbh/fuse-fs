package pkg_test

import (
	"context"
	"fmt"
	"fuse-filesystem/pkg"
	"testing"

	"bazil.org/fuse"
	"github.com/stretchr/testify/assert"
)

type MockINodeGenerator struct{}

func (MockINodeGenerator) Next() uint64 {
	return 1
}

func TestDirectoryHasCorrectPermissionAttr(t *testing.T) {
	dir := pkg.NewDir(0, MockINodeGenerator{})
	attr := fuse.Attr{}
	err := dir.Attr(context.TODO(), &attr)
	assert.NoError(t, err)
	assert.Equal(t,
		attr.Mode&pkg.DirectoryPermission,
		pkg.DirectoryPermission,
		fmt.Sprintf("The directory should have %d permissions", pkg.DirectoryPermission))
}

func TestDirectoryHasInitialBlockSizeofDefaultSize(t *testing.T) {
	dir := pkg.NewDir(0, MockINodeGenerator{})
	attr := fuse.Attr{}
	err := dir.Attr(context.TODO(), &attr)
	assert.NoError(t, err)
	assert.Equal(t,
		attr.Size,
		pkg.DefaultSize,
		fmt.Sprintf("The directory should have size of %d", pkg.DefaultSize))
}

func TestDirectoryShouldBeAbleToListItself(t *testing.T) {
	dirs, err := pkg.NewDir(0, MockINodeGenerator{}).ReadDirAll(context.TODO())
	assert.NoError(t, err)

	var selfDir *fuse.Dirent
	for _, dir := range dirs {
		if dir.Name == "." {
			selfDir = &dir
		}
	}

	assert.NotNilf(t, selfDir, "the directory should be able to list itself")
}

func TestDirectoryShouldBeAbleToListItsParent(t *testing.T) {
	dirs, err := pkg.NewDir(0, MockINodeGenerator{}).ReadDirAll(context.TODO())
	assert.NoError(t, err)

	var parentDir *fuse.Dirent
	for _, dir := range dirs {
		if dir.Name == ".." {
			parentDir = &dir
		}
	}

	assert.NotNilf(t, parentDir, "the directory should be able to list its parent")
}

func TestRootDirectoryShouldHaveItselfAsParent(t *testing.T) {
	dirs, err := pkg.NewDir(0, MockINodeGenerator{}).ReadDirAll(context.TODO())
	assert.NoError(t, err)

	var selfDir *fuse.Dirent
	var parentDir *fuse.Dirent
	for _, dir := range dirs {
		if dir.Name == "." {
			selfDir = &dir
		}
		if dir.Name == ".." {
			parentDir = &dir
		}
	}

	assert.Equalf(t, selfDir, parentDir, "if the directory is root, its parent should be itself")
}
