package main

import "testing"

func Test1(t *testing.T) {
	input := 121
	output := isPalindrome(input)

	want := true

	if output != want {
		t.Errorf("Want %v but got %v", want, output)
	}
}

func Test2(t *testing.T) {
	input := -121
	output := isPalindrome(input)

	want := false

	if output != want {
		t.Errorf("Want %v but got %v", want, output)
	}
}
func Test3(t *testing.T) {
	input := 10
	output := isPalindrome(input)

	want := false

	if output != want {
		t.Errorf("Want %v but got %v", want, output)
	}
}
