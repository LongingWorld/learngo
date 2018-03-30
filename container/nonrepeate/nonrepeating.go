package nonrepeate

import (
	"fmt"
)

func lengthOfNonRepeatingSubStr(s string) (int, string) {
	lastOccurred := make(map[rune]int)
	start := 0 //the first
	finastart := 0
	maxlength := 0
	sliceS := []rune(s)

	for i, ch := range sliceS { //travering sliceS,get one character of the input parameter
		//fmt.Printf("the %d character of sliceS is %c\n",i,ch)
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
			//fmt.Printf("the value of the %c " +
			//	"in the lastOccurred is %d\n",ch,lastI)
		}

		if i-start+1 > maxlength {
			maxlength = i - start + 1
			finastart = i - maxlength + 1
		}
		lastOccurred[ch] = i

	}
	//fmt.Printf("start :%d ,maxlength : %d \n",start,maxlength)
	//fmt.Printf("finastart :%d ,maxlength : %d \n",finastart,maxlength)
	//fmt.Println(sliceS)
	sliceS = sliceS[finastart : finastart+maxlength]

	fmt.Println(lastOccurred)

	return maxlength, string(sliceS)
}

func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("abcdddssafadf"))
	fmt.Println(lengthOfNonRepeatingSubStr(""))
	fmt.Println(lengthOfNonRepeatingSubStr("abcdefga"))
	fmt.Println(lengthOfNonRepeatingSubStr("bbbbbbbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("厉害了我的国"))
}
