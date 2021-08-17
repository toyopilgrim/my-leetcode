package atoi

import "testing"

func Test1(t *testing.T) {
	input := "-111"

	output := myAtoi(input)

	want := -111
	if output != want {
		t.Errorf("Want %v but got %v", want, output)
	}
}

func Test2(t *testing.T) {
	input := "-91283472332"

	output := myAtoi(input)

	want := -2147483648
	if output != want {
		t.Errorf("Want %v but got %v", want, output)
	}
}

func Test3(t *testing.T) {
	input := "91283472332"

	output := myAtoi(input)

	want := 2147483647
	if output != want {
		t.Errorf("Want %v but got %v", want, output)
	}
}

func Test4(t *testing.T) {
	input := "  +42"

	output := myAtoi(input)

	want := 42
	if output != want {
		t.Errorf("Want %v but got %v", want, output)
	}
}

func Test5(t *testing.T) {
	input := "  + 42"

	output := myAtoi(input)

	want := 0
	if output != want {
		t.Errorf("Want %v but got %v", want, output)
	}
}

func Test6(t *testing.T) {
	input := "-"

	output := myAtoi(input)

	want := 0
	if output != want {
		t.Errorf("Want %v but got %v", want, output)
	}
}
