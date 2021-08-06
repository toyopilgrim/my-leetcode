package median_two_sorted_arrays

import "testing"

func TestWithOdd(t *testing.T) {
	input1 := []int{1, 2}
	input2 := []int{3, 4, 5}

	output := findMedianSortedArrays(input1, input2)

	want := 3.0

	if output != want {
		t.Errorf("Want %v but got %v", want, output)
	}
}

func TestWithEven(t *testing.T) {
	input1 := []int{1, 2}
	input2 := []int{3, 4}

	output := findMedianSortedArrays(input1, input2)

	want := 2.5

	if output != want {
		t.Errorf("Want %v but got %v", want, output)
	}
}
