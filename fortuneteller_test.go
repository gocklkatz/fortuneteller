package fortuneteller

import "testing"

func TestTellFortune(t *testing.T) {
	want := "The early bird catches the worm!"
	if got := TellFortune(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}