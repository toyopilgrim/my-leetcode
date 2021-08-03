package two_nums

import (
	"math/big"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	strNum1 := convertListNodeToStrNum(l1)
	strNum2 := convertListNodeToStrNum(l2)

	sum := toBigInt(strNum1).Add(toBigInt(strNum1), toBigInt(strNum2))
	sumStr := sum.String()

	newListNode := &ListNode{}
	setReversedListNode(newListNode, sumStr, len(sumStr)-1)
	return newListNode
}

func convertListNodeToStrNum(l *ListNode) string {
	var nums []int
	numsList := appendNums(l, nums)
	var strNum string
	for i := len(numsList) - 1; i >= 0; i-- {
		strNum += strconv.Itoa(numsList[i])
	}
	return strNum
}

func setReversedListNode(l *ListNode, sumStr string, index int) *ListNode {
	l.Val = toInt(string(sumStr[index]))
	if index == 0 {
		return l
	}
	l.Next = &ListNode{}
	return setReversedListNode(l.Next, sumStr, index-1)
}

func appendNums(l *ListNode, res []int) []int {
	res = append(res, l.Val)
	if l.Next == nil {
		return res
	}
	return appendNums(l.Next, res)
}

func toBigInt(s string) *big.Int {
	var num, _ = new(big.Int).SetString(s, 10)
	return num
}

func toInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}
