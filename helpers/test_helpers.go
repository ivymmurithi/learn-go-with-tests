package helpers

import (
	"testing"
)

func AssertCorrectMessage(t testing.TB, got, want interface{}) {
	t.Helper()

	switch gotTyped := got.(type) {
	case string:
		if wantTyped, ok := want.(string); ok && gotTyped != wantTyped {
			t.Errorf("got %q, want %q", gotTyped, wantTyped)
		}
	case int:
		if wantTyped, ok := want.(int); ok && gotTyped != wantTyped {
			t.Errorf("expected %d got %d", wantTyped, gotTyped)
		}
	default:
		t.Errorf("unsupported type: got %T, want %T", got, want)
	}
}
