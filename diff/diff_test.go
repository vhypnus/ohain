package diff

import (
	"testing"
	"fmt"
)



func TestNewDiff(t *testing.T) {
	var content = `diff --git a/check/check.go b/check/check.go
	new file mode 100644
	index 0000000..687b9b2
	--- /dev/null
	+++ b/check/check.go
	@@ -0,0 +1,20 @@
	+package check
	+
	+
	+
	+type check interface {
	+       check(f *File) string
	+}

	diff --git a/model/model.go b/model/model.go
	deleted file mode 100644
	index 80b2c5b..0000000
	--- a/model/model.go
	+++ /dev/null
	@@ -1,15 +0,0 @@
	-package model
	-
	-type Model struct {
	-
	-       note string
	-
	-       type string
	-
	-       variable string
	-}
	-
	-
	-func NewModel(path string) []Model {
	-
	-}`

	var result = NewDiff(content)

	fmt.Println(result)
}

