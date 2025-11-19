package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat 5 times", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("repeat 0 times", func(t *testing.T) {
		got := Repeat("a", 0)
		want := ""

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	// ... setup ...
	for b.Loop() {
		// ... code to measure ...
		Repeat("a", 10)
	}
	// ... cleanup ...
}

func ExampleRepeat() {
	repeated := Repeat("z", 3)
	fmt.Println(repeated)
	// Output: zzz
}
