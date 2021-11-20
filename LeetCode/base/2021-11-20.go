package main

import "fmt"

func main() {
	val := twoSum([]int{1, 2, 3, 4, 5}, 5)
	fmt.Printf("%v", val)
}

/*
1. 两数之和
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target 的那两个整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
你可以按任意顺序返回答案
*/

func twoSum(nums []int, target int) []int {
	var result []int
	var mapResult = make(map[int]int, 2)
	for i, v := range nums {
		sub := target - v
		if j, ok := mapResult[sub]; ok {
			result = append(result, i)
			result = append(result, j)
			return result
		} else {
			mapResult[v] = i
		}
	}

	return result
}
