package main

import (
	"testing"
)

func TestNonrepeatingSubStr(t *testing.T) {
	tests := []struct {
		s string
		i int
	}{
		//normal cases
		{"adbasdfasf",5},
		{"acbdfa",5},


		//edge cases
		{"aaaa",1},
		{"",0},
		{"b",1},
		{"abcabcabcd",4},

		//Chinese support
		{"厉害了我的国",6},
		{"黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花",8},

	}

	for _,tt := range tests  {
		actual := LengthoOfNoneRepeatSubStr(tt.s)
		if actual != tt.i{
			t.Errorf("got %d for input %s;" + "expected %d ",actual,tt.s,tt.i)
		}
	}
}

func BenchmarkSubStr(b *testing.B)  {
	s := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"
	for i := 0; i < 13  ; i++ {
		s = s + s
	}
	n := 8

	b.Logf("len(s) = %d",len(s))
	b.ResetTimer()   //重置时钟

	for i:=0;i < b.N ; i++ {
		actual := LengthoOfNoneRepeatSubStr(s)
		if actual != n{
			b.Errorf("got %d for input %s;" + "expected %d ",actual,s,n)
		}
	}
}