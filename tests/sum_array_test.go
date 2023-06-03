package tests

import (
	"testing"
	"reflect"
 	"github.com/rixtrayker/learn-go-tests/arrays"
)

func TestSum(t *testing.T) {
	deepCheckSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	checkSums := func(t testing.TB, got, want int, numbers []int) {
		t.Helper()
		if got != want {
			t.Errorf("want '%d' but got '%d', given %v", want, got, numbers)
		}
	}
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := arrays.Sum(numbers)

		want := 15

		checkSums(t, got, want, numbers)
	})

	t.Run("collection of any numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := arrays.Sum(numbers)
		want := 6

		checkSums(t, got, want, numbers)
	})

	t.Run("deep equal", func(t *testing.T) {

		got := arrays.SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}
	
		deepCheckSums(t, got, want)
	})

	t.Run("make the sums of some slices tails", func(t *testing.T) {
			got := arrays.SumAllTails([]int{1, 2}, []int{0, 9})
			want := []int{2, 9}
	
			deepCheckSums(t, got, want)
		})

	t.Run("safely sum empty slices tails", func(t *testing.T) {
		got := arrays.SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		deepCheckSums(t, got, want)
	})
}