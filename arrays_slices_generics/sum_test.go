package arraysslices

import (
	"reflect"
	"strings"
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

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got []int, want []int) {
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

func TestReduce(t *testing.T) {
	t.Run("multiplication of all elements", func(t *testing.T) {
		numbers := []int{4, 5, 10}

		multiply := func(result int, value int) int {
			return result * value
		}

		got := Reduce(numbers, multiply, 1)
		want := 200

		AssertEqual(t, got, want)
	})

	t.Run("concatenate strings", func(t *testing.T) {
		strs := []string{"Hello, ", "world", "!"}

		concatenate := func(result string, value string) string {
			return result + value
		}

		got := Reduce(strs, concatenate, "")
		want := "Hello, world!"

		AssertEqual(t, got, want)
	})
}

func TestBadBank(t *testing.T) {
	var (
		riya  = Account{Name: "Riya", Balance: 100}
		chris = Account{Name: "Chris", Balance: 75}
		adil  = Account{Name: "Adil", Balance: 200}

		transactions = []Transaction{
			NewTransaction(chris, riya, 100),
			NewTransaction(adil, chris, 25),
		}
	)

	newBalanceFor := func(account Account) float64 {
		return NewBalanceFor(account, transactions).Balance
	}

	AssertEqual(t, newBalanceFor(riya), 200)
	AssertEqual(t, newBalanceFor(chris), 0)
	AssertEqual(t, newBalanceFor(adil), 175)
}

func TestFind(t *testing.T) {
	t.Run("find first even number", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		firstEvenNumber, found := Find(numbers, func(x int) bool {
			return x%2 == 0
		})
		AssertTrue(t, found)
		AssertEqual(t, firstEvenNumber, 2)
	})

	type Person struct {
		Name string
	}
	
	t.Run("Find the best programmer", func(t *testing.T) {
		people := []Person{
			Person{Name: "Kent Beck"},
			Person{Name: "Martin Fowler"},
			Person{Name: "Chris James"},
		}
	
		king, found := Find(people, func(p Person) bool {
			return strings.Contains(p.Name, "Chris")
		})
	
		AssertTrue(t, found)
		AssertEqual(t, king, Person{Name: "Chris James"})
	})
}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()

	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func AssertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()

	if got == want {
		t.Errorf("didn't want %+v", got)
	}
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()

	if !got {
		t.Errorf("got %v, wanted true", got)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()

	if got {
		t.Errorf("got %v, wanted false", got)
	}
}
