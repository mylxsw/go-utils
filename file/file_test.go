package file_test

import (
	"testing"

	"github.com/mylxsw/go-utils/assert"
	"github.com/mylxsw/go-utils/file"
)

func TestLocalFS_ListFiles(t *testing.T) {
	fs := file.LocalFS{}
	files, err := fs.ListFiles(".")
	assert.NoError(t, err)
	assert.True(t, len(files) > 0)
}
