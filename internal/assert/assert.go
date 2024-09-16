package assert

import "testing"

func Equal[T comparable](t *testing.T, got, want T) {
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func NotEqual[T comparable](t *testing.T, got, want T) {
	if got == want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Error(t *testing.T, got error) {
	if got == nil {
		t.Errorf("got nil, want error")
	}
}

func NotError(t *testing.T, got error) {
	if got != nil {
		t.Errorf("got %v, want nil", got)
	}
}
