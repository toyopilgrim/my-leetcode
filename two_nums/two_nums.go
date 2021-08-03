package two_nums

import (
	"math/big"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var numsL1 []int
	numsList1 := appendNums(l1, numsL1)
	var strNum1 string
	for i := len(numsList1) - 1; i >= 0; i-- {
		strNum1 += strconv.Itoa(numsList1[i])
	}

	var numsL2 []int
	numsList2 := appendNums(l2, numsL2)
	var strNum2 string
	for i := len(numsList2) - 1; i >= 0; i-- {
		strNum2 += strconv.Itoa(numsList2[i])
	}

	sum := toBigInt(strNum1).Add(toBigInt(strNum1), toBigInt(strNum2))
	sumStr := sum.String()

	newListNode := &ListNode{}
	cnt := len(sumStr)

	getRes(newListNode, sumStr, cnt-1)
	return newListNode
}

func getRes(l *ListNode, sumStr string, index int) *ListNode {
	l.Val = toInt(string(sumStr[index]))
	if index == 0 {
		return l
	}
	l.Next = &ListNode{}
	return getRes(l.Next, sumStr, index-1)
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
