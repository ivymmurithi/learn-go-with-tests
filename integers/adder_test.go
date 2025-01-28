package integers

import (
	"fmt"
	"testing"
	"github.com/ivymmurithi/learn-go-with-tests/helpers"
)

func TestAdd(t *testing.T) {
	got := Add(2, 2)
	want := 4

	helpers.AssertCorrectMessage(t, got, want)
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	//Output: 6
}