package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
	got := Hello("AB")
	want := "Namaste, AB"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
})
}