package main

import "fmt"

var nums  =[]int{0,0,1,1,2,2,3,3,4,4,5,5,6,6,7,7,8,8,9,9,10,10}

func main() {
	count := removeDuplicates(nums)
	fmt.Println("length of slice is ",count)
	for i:=0;i < count;i++{
		fmt.Printf("%d ",nums[i])
	}
	//fmt.Println()
}

func removeDuplicates(nums []int) int {
	size := 0
	for _,v := range nums{
		if v != nums[size]{
			size+=1
			nums[size] =v
		}
	}
	nums = nums[0:size+1]
	return size+1
}
