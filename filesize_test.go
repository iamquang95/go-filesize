package filesize

import (
	"testing"
)

func TestConvert(t *testing.T) {
	var x = Byte(100)
	var converted = x.Convert(KB)
	var expected = float64(100) / float64(1024)
	t.Log(converted)
	t.Log(expected)
	if converted != expected {
		t.Error("failed")
	}
}
