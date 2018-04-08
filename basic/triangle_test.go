package basic_test

import (
	"GitHub/learngo/basic"
	"testing"
)

func TestTriangle(t *testing.T) {
	tests := []struct{a,b,c int}{
		{3,4,5},
		{5,12,13},
		{8,15,17},
		{12,35,37},
		{30000,40000,50000},
	}

	for _,value := range tests {
		if actual := basic.CalcTriangle(value.a,value.b);actual != value.c{
			t.Errorf("CalcTriangle(%d, %d); "+ " got %d ;" + "expected %d",value.a,value.b,actual,value.c)
		}
	}
}
