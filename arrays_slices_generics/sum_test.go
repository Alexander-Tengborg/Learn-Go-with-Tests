package arraysslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given %d", got, want, numbers)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3,9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got[]int, want[]int) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("gets the tail slice of 2 slices of length 2", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		
		checkSums(t, got, want)
	})

	t.Run("gets the tail slice of 2 slices of length 3", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 5}, []int{0, 9, 3})
		want := []int{7, 12}
		
		checkSums(t, got, want)

	})

	t.Run("gets the tail slice of 1 slice of length 0, and 1 of length 2", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		want := []int{0, 9}
		
		checkSums(t, got, want)
	})
}