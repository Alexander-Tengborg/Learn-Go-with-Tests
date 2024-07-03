package walk

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	t.Run("All sorts of tests", func(t *testing.T) {
		cases := []struct {
			Name          string
			Input         any
			ExpectedCalls []string
		}{
			{
				"struct with one string field",
				struct {
					Name string
				}{"Linda"},
				[]string{"Linda"},
			},
			{
				"struct with two string fields",
				struct {
					Name string
					City string
				}{"Linda", "Tokyo"},
				[]string{"Linda", "Tokyo"},
			},
			{
				"struct with one string field and one int field",
				struct {
					Name string
					Age  int
				}{"Linda", 23},
				[]string{"Linda"},
			},
			{
				"struct with nested structs",
				Person{"Linda",
					Profile{23, "Tokyo"},
				},
				[]string{"Linda", "Tokyo"},
			},
			{
				"pointer to a struct",
				&Person{"Linda",
					Profile{23, "Tokyo"},
				},
				[]string{"Linda", "Tokyo"},
			},
			{
				"slices",
				[]Profile{
					{21, "Paris"},
					{93, "New York"},
				},
				[]string{"Paris", "New York"},
			},
			{
				"arrays",
				[2]Profile{
					{21, "Paris"},
					{93, "New York"},
				},
				[]string{"Paris", "New York"},
			},
		}

		for _, test := range cases {
			t.Run(test.Name, func(t *testing.T) {
				var got []string
				walk(test.Input, func(input string) {
					got = append(got, input)
				})

				if !reflect.DeepEqual(got, test.ExpectedCalls) {
					t.Errorf("got %q, want %q", got, test.ExpectedCalls)
				}
			})
		}
	})

	t.Run("maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Bäääh",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Bäääh")
	})

	t.Run("channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Dortmund"}
			aChannel <- Profile{23, "Richmond"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Dortmund", "Richmond"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("functions", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Dortmund"}, Profile{23, "Richmond"}
		}

		var got []string
		want := []string{"Dortmund", "Richmond"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, got []string, want string) {
	t.Helper()
	contains := false

	for _, x := range got {
		if x == want {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", got, want)
	}
}
