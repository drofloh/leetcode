package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	ret := make([]int, 2)

	for i := 0; i < len(nums); i++ {
		if nums[i]+nums[len(nums)-1-i] == target {
			ret[0], ret[1] = i, len(nums)-1-i
			break
		}
	}

	fmt.Println(ret)

}
