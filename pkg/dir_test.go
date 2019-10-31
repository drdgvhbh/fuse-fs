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
