package diff_test

import (
	"testing"

	"github.com/mylxsw/go-utils/assert"
	"github.com/mylxsw/go-utils/diff"
	"github.com/mylxsw/go-utils/file"
)

func TestNewDiffer(t *testing.T) {
	differ := diff.NewDiffer(file.LocalFS{}, "/tmp", 0)
	assert.Equal(t, `--- s1
+++ s2
@@ -2 +2 @@
-Are you ok?
+What's your name?
`, differ.Diff("s1", "Hello, world\nAre you ok?", "s2", "Hello, world\nWhat's your name?"))
}
