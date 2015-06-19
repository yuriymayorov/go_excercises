package strutil

import "testing"

func TestIsUniqueChars(t *testing.T) {
	cases := []struct {
		in string
		want bool
	}{
		{"Hello world", false},
		{"Hola", true},
	}	
	for _, c := range cases {
		got := IsUniqueChars(c.in)
		if got != c.want {
			t.Errorf("IsUniqueChars(%q) == %t, want %t", c.in, got, c.want)
		}
	}
}