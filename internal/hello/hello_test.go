package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people in english", func(t *testing.T) {
		got := Hello("AB","English")
		want := "Hello, AB"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("saying hello to the world", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, world"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("saying hello to people in hindi", func(t *testing.T) {
		got := Hello("AB", "Hindi")
		want := "Namaste, AB"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	
}
