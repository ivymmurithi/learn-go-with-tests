package sumoflsit

import (
	"testing"
	"reflect"
	"github.com/ivymmurithi/learn-go-with-tests/helpers"
)

func TestSumOfList(t *testing.T) {
	t.Run("summation of 4 numbers", func(t *testing.T) {
		numbersList := [4]int{1, 3, 4, 6}
		got := GenerateSum(numbersList)
		want := 14

		helpers.AssertCorrectMessage(t, got, want)
	})

	t.Run("summation of dynamic list", func(t *testing.T) {
		dynamicNumbersList :=[]int{4, 5, 10, 20, 10}
		got := GenerateDynamicListSum(dynamicNumbersList)
		want := 49

		helpers.AssertCorrectMessage(t, got, want)
	})
	
}

func TestSumAl(t *testing.T) {
	got := SumAll([]int{1,2,4}, []int{5,5})
	want := []int{7, 10}

	// not equality operator for slices
	// deep equal is not type safe eg want could be "Bob" but it will compile as usual
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}