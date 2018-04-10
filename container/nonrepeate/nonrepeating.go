package main//nonrepeate

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

	//fmt.Println(lastOccurred)
	//fmt.Println("Traversing Map:") //map是无序的
	//for k, v := range lastOccurred {
	//	fmt.Printf("key : %c, value : %d \n",k, v)
	//}

	return maxlength, string(sliceS)
}

var lastOccurred = make([]int,0xffff)  //存放不同字符出现的最后位置

func LengthoOfNoneRepeatSubStr(s string) int  {
	//lastOccurred := make([]int,0xffff)   //65535 内存大小   空间换时间
	for i := range lastOccurred  {
		lastOccurred[i] = -1
	}
	//lastOccurred := make(map[rune]int)  //存放不同字符出现的最后位置
	start,maxlength := 0,0   //start 记录无重复子串的起始位置  maxlength记录最长无重复子串的长度

	for i,ch :=range []rune(s) {   //rune
		if lastI := lastOccurred[ch]; lastI != -1 && lastI >= start  {  //每个字符重复出现后，更新start为此字符出现的位置的后一位
			start = lastI +1
		}
		if maxlength < i-start + 1 {   //由0...i...到len(s)遍历字符串s 最大长度为(i - start +1 )
			 maxlength = i - start + 1
		}
		lastOccurred[ch] = i   //记录不同字符出现的最后位置
	}
	return maxlength
}



func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("abcdddssafadf"))
	fmt.Println(lengthOfNonRepeatingSubStr(""))
	fmt.Println(lengthOfNonRepeatingSubStr("abcdefga"))
	fmt.Println(lengthOfNonRepeatingSubStr("bbbbbbbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("厉害了我的国"))
}
