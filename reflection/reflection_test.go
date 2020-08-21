package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Nahua"},
			[]string{"Nahua"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Nahua", "Berlin"},
			[]string{"Nahua", "Berlin"},
		},
		{
			"Struct with non-string field",
			struct {
				Name string
				Age  int
			}{"Nahua", 29},
			[]string{"Nahua"},
		},
		{
			"Nested fields",
			Person{
				"Nahua",
				Profile{29, "Berlin"},
			},
			[]string{"Nahua", "Berlin"},
		},
		{
			"Pointers to things",
			&Person{
				"Nahua",
				Profile{29, "Berlin"},
			},
			[]string{"Nahua", "Berlin"},
		},
		{
			"Slices",
			[]Profile{
				{29, "Berlin"},
				{34, "Reykjavik"},
			},
			[]string{"Berlin", "Reykjavik"},
		},
		{
			"Arrays",
			[2]Profile{
				{29, "Berlin"},
				{34, "Reykjavik"},
			},
			[]string{"Berlin", "Reykjavik"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v want %v", got, test.ExpectedCalls)
			}
		})
	}
}
