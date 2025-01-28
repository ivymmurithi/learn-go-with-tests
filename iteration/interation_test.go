package iteration

import (
	"fmt"
	"testing"
	"github.com/ivymmurithi/learn-go-with-tests/helpers"
)

func TestIteration(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"

		helpers.AssertCorrectMessage(t, got, want)
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeat := Repeat("i", 10)
	fmt.Println(repeat)
	// Output: iiiiiiiiii
}