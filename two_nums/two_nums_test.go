package two_nums

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	input1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	input2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	got := AddTwoNumbers(input1, input2)

	if got.Val != 7 {
		t.Errorf("got %d want %d ", got.Val, 7)
	}
	if got.Next.Val != 0 {
		t.Errorf("got %d want %d ", got.Val, 0)
	}
	if got.Next.Next.Val != 8 {
		t.Errorf("got %d want %d ", got.Val, 8)
	}
	if got.Next.Next.Next != nil {
		t.Errorf("Val wan't nil")
	}

}
