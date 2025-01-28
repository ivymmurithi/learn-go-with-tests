package sumoflsit

import (
	"testing"
	"slices"
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

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()

		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
		
	}
	t.Run("Summation of tails", func(t *testing.T) {
		got := SumAllTails([]int{1,2,6}, []int{5,10})
		want := []int{8, 10}

		checkSums(t, got, want)
	})

	t.Run("safely summate an empty list", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{4,6,2})
		want := []int{0, 8}

		checkSums(t, got, want)
	})
	
}