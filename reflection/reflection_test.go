package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name string
	Profile
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
			"struct with one string field",
			struct {
				Name string
			}{
				"Marco",
			},
			[]string{"Marco"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{
				"Marco",
				"Yangon",
			},
			[]string{"Marco", "Yangon"},
		},
		{
			"struct with non-string field",
			struct {
				Name string
				Age  int
			}{
				"Marco",
				23,
			},
			[]string{"Marco"},
		},
		{
			"nested fields",
			Person{
				"Marco",
				Profile{
					23,
					"Yangon",
				},
			},
			[]string{"Marco", "Yangon"},
		},
		{
			"pointers to things",
			&Person{
				"Marco",
				Profile{
					23,
					"Yangon",
				},
			},
			[]string{"Marco", "Yangon"},
		},
		{
			"slices",
			[]Profile{
				{23, "Yangon"},
				{22, "Tsukuba"},
			},
			[]string{"Yangon", "Tsukuba"},
		},
		{
			"arrays",
			[2]Profile{
				{23, "Yangon"},
				{22, "Tsukuba"},
			},
			[]string{"Yangon", "Tsukuba"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string

			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("wrong number of function calls got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}
		got := []string{}
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{23, "Yangon"}
			aChannel <- Profile{22, "Tsukuba"}
			close(aChannel)
		}()

		got := []string{}
		want := []string{"Yangon", "Tsukuba"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{23, "Yangon"}, Profile{22, "Tsukuba"}
		}
		got := []string{}
		want := []string{"Yangon", "Tsukuba"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false

	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %v to contain %q, but didn't", haystack, needle)
	}
}
