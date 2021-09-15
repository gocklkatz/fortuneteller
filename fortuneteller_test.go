package fortuneteller

import "testing"

func TestTellFortune(t *testing.T) {
	want := "The early bid catches something!"
	if got := tellFortune(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}