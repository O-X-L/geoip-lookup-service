package u

import (
	"testing"
)

func TestGetMapValue(t *testing.T) {
	data := map[string]interface{}{
		"country": "Austria",
		"details": map[string]interface{}{
			"city": "Vienna",
		},
	}

	val := GetMapValue(data, "country")
	if val != "Austria" {
		t.Errorf("Expected Austria, got %v", val)
	}
}
