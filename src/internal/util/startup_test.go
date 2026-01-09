package u

import (
	"testing"
)

func TestPrependDBDirPath(t *testing.T) {
	dir := "/tmp/db"
	file := "test.mmdb"
	expected := "/tmp/db/test.mmdb"

	result := prependDBDirPath(dir, file)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}

	// Test with absolute path (should not prepend)
	absFile := "/etc/db/geo.mmdb"
	resultAbs := prependDBDirPath(dir, absFile)
	if resultAbs != absFile {
		t.Errorf("Expected %s, got %s", absFile, resultAbs)
	}
}
