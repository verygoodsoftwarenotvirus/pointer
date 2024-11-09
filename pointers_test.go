package pointer

import (
	"testing"
)

func TestTo(T *testing.T) {
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

func TestForSlice(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		expected := []bool{false}
		actual := ForSlice(expected)

		if actual == nil {
			t.Fatalf("expected %v, got nil", expected)
		}

		if len(actual) != len(expected) {
			t.Fatalf("expected %d items, got %d", len(expected), len(actual))
		}

		for i, x := range expected {
			if x != *actual[i] {
				t.Fatalf("expected %v, got %v", expected, *actual[i])
			}
		}
	})
}

func TestDeref(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		rawExpected := "things"
		expected := &rawExpected
		actual := Deref(expected)

		if actual != *expected {
			t.Fatalf("expected %v, got %v", expected, actual)
		}
	})
}

func TestDerefSlice(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		rawExpected := []string{"things"}
		expected := []*string{&rawExpected[0]}
		actual := DerefSlice(expected)

		for i, x := range expected {
			if *x != actual[i] {
				t.Fatalf("expected %v, got %v", expected, &actual[i])
			}
		}
	})
}

func TestForMap(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		expected := map[string]bool{"things": false}
		actual := ForMap(expected)

		if actual == nil {
			t.Fatalf("expected %v, got nil", expected)
		}

		if len(actual) != len(expected) {
			t.Fatalf("expected %d items, got %d", len(expected), len(actual))
		}

		for k, v := range expected {
			if v != *actual[k] {
				t.Fatalf("expected %v, got %v", expected, *actual[k])
			}
		}
	})
}

func TestDerefMap(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		exampleInput := map[string]*string{"test": To("things")}
		expected := map[string]string{"test": "things"}
		actual := DerefMap(exampleInput)

		for k, v := range expected {
			if actual[k] != v {
				t.Fatalf("expected %v, got %v", expected, v)
			}
		}
	})
}
