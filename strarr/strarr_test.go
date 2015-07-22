package strarr

import (
	"testing"
	"reflect"
)

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

func TestRotate(t *testing.T) {
	cases := []struct {
		in1 [][]int
		want [][]int
	}{
		{[][]int{{1,2},{3,4}}, [][]int{{3,1}, {4,2}}},
		{[][]int{{1,2,3},{4,5,6},{7,8,9}}, [][]int{{7,4,1},{8,5,2},{9,6,3}}},
	}	
	for _, c := range cases {
		got := Rotate(c.in1)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Rotate(%v) == %v, want %v", c.in1, got, c.want)
		}
	}
}

func TestSetZeros(t *testing.T) {
	cases := []struct {
		in1 [][]int
		want [][]int
	}{
		{[][]int{{1,2,3},{4,0,5}}, [][]int{{1,0,3}, {0,0,0}}},
		{[][]int{{1,2,3},{0,5,0},{7,8,9}}, [][]int{{0,2,0},{0,0,0},{0,8,0}}},
	}	
	for _, c := range cases {
		got := SetZeros(c.in1)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Rotate(%v) == %v, want %v", c.in1, got, c.want)
		}
	}
}

func TestIsRotation(t *testing.T) {
	cases := []struct {
		in1 string
		in2 string
		want bool
	}{
		{"waterbottle", "erbottlewat", true},
		{"football", "ballfoot", true},
		{"computer", "trcompu", false},
	}	
	for _, c := range cases {
		got := IsRotation(c.in1, c.in2)
		if got != c.want {
			t.Errorf("IsRotation(%s, %s) == %v, want %v", c.in1, c.in2, got, c.want)
		}
	}		
}