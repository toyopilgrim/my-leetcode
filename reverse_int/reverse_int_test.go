package reverse_int

import "testing"

func Test1(t *testing.T) {
	input := 123

	output := reverse(input)
	expected := 321

	if output != expected {
		t.Errorf("Want %v but got %v", expected, output)
	}
}

func Test2(t *testing.T) {
	input := -123

	output := reverse(input)
	expected := -321

	if output != expected {
		t.Errorf("Want %v but got %v", expected, output)
	}
}
