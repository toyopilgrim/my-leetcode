package two_sum

import (
	"reflect"
	"testing"
)

func TestTwoSum1(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9

	got := TwoSum(nums, target)
	want := []int{0, 1}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d want %d given, nums: %v and target: %v", got, want, nums, target)
	}
}

func TestTwoSum2(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6

	got := TwoSum(nums, target)
	want := []int{1, 2}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d want %d given, nums: %v and target: %v", got, want, nums, target)
	}
}

func TestTwoSum3(t *testing.T) {
	nums := []int{3, 3}
	target := 6

	got := TwoSum(nums, target)
	want := []int{0, 1}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d want %d given, nums: %v and target: %v", got, want, nums, target)
	}
}
