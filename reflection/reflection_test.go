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

	t.Run("with Maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

	t.Run("with Channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{29, "Berlin"}
			aChannel <- Profile{34, "Reykjavik"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Berlin", "Reykjavik"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

  t.Run("with function", func(t *testing.T) {
    aFunction := func() (Profile, Profile) {
      return Profile{29, "Berlin"}, Profile{34, "Reykjavik"}
    }

    var got []string
    want := []string{"Berlin", "Reykjavik"}

    walk(aFunction, func(input string) {
      got = append(got, input)
    })

    if !reflect.DeepEqual(got, want) {
      t.Errorf("got %v want %v", got, want)
    }
  })
}

func assertContains(t *testing.T, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
