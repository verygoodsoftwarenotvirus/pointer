package pointer

import (
	"testing"
)

func TestPointer(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		expected := false
		actual := To(expected)

		if actual == nil {
			t.Fatalf("expected %v, got nil", expected)
		}

		if *actual != expected {
			t.Fatalf("expected %v, got %v", expected, *actual)
		}
	})
}
