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

func TestReverse(t *testing.T) {
	cases := []struct {
		in string
		want string
	}{
		{"holla", "alloh"},
		{"pass", "ssap"},
		{"", ""},
	}	
	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%s) == %s, want %s", c.in, got, c.want)
		}
	}
}

func TestIsPermutation(t *testing.T) {
	cases := []struct {
		in1 string
		in2 string
		want bool
	}{
		{"holla", "alloh", true},
		{"red", "ret", false},
		{"game", "emag", true},
	}	
	for _, c := range cases {
		got := IsPermutation(c.in1, c.in2)
		if got != c.want {
			t.Errorf("IsPermutation(%s, %s) == %s, want %s", c.in1, c.in2, got, c.want)
		}
	}
}

func TestCompress(t *testing.T) {
	cases := []struct {
		in1 string
		want string
	}{
		{"aaaab", "a4b1"},
		{"aa", "aa"},
		{"aab", "aab"},
		{"footballllllllllll", "f1o2t1b1a1l12"},
	}	
	for _, c := range cases {
		got := Compress(c.in1)
		if got != c.want {
			t.Errorf("Compress(%s) == %s, want %s", c.in1, got, c.want)
		}
	}
}