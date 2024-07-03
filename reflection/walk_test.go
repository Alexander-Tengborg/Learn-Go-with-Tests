package walk

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name string
		Input any
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
				Age int
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


}