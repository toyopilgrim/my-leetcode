package zigza_conv

import "testing"

func Test1(t *testing.T){
	s := "PAYPALISHIRING"
	numRows := 4

	output := convert(s,numRows)

	want := "PINALSIGYAHRPI"
	if output != want{
		t.Errorf("Want %v but got %v",want, output)
	}
}

func Test2(t *testing.T){
	s := "PAYPALISHIRING"
	numRows := 3

	output := convert(s,numRows)

	want := "PAHNAPLSIIGYIR"
	if output != want{
		t.Errorf("Want %v but got %v",want, output)
	}
}

func Test3(t *testing.T){
	s := "AB"
	numRows := 1

	output := convert(s,numRows)

	want := "AB"
	if output != want{
		t.Errorf("Want %v but got %v",want, output)
	}
}