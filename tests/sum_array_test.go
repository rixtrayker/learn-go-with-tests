package tests

import (
	"testing"
	"reflect"
 	"github.com/rixtrayker/learn-go-tests/arrays"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := arrays.Sum(numbers)

		want := 15

		if got != want {
			t.Errorf("want '%d' but got '%d', given %v", want, got, numbers)
		}
	})

	t.Run("collection of any numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := arrays.Sum(numbers)

		want := 6

		if got != want {
			t.Errorf("want '%d' but got '%d', given %v", want, got, numbers)
		}
	})

	t.Run("deep equal", func(t *testing.T) {

		got := arrays.SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}
	
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
