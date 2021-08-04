package substr_wo_repeat

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	input := "pwwkew"

	output := LengthOfLongestSubstring(input)
	want := 3

	if output != want {
		t.Errorf("Want %d but got %d", want, output)
	}
}
