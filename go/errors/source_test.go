package errors

import "testing"

func TestCaptureSource(t *testing.T) {
	s := CaptureSource(1)
	if s == nil {
		t.Error("CaptureSource should not return nil")
	}
}
