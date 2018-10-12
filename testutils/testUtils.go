package testutils

import "testing"

func ASSERT_EQ(t *testing.T, actual interface{}, expected interface{}) {
	if expected != actual {
		t.Fatalf("expected=%d, got=%d", expected, actual)
	}
}
