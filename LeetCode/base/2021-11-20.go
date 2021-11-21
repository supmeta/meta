package main

import "fmt"

func main() {
	val := twoSum([]int{1, 2, 3, 4, 5}, 5)
	fmt.Printf("%v", val)

	val2 := reverse(2147483647)
	fmt.Printf("%v", val2)
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

/*
7. 整数反转
给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
如果反转后整数超过 32 位的有符号整数的范围[−2^31, 2^31− 1] ，就返回 0。
假设环境不允许存储 64 位整数（有符号或无符号）。
提示：-231 <= x <= 231 - 1
*/
func reverse(x int) int {
	var result = 0
	max := 1 << 31
	if -max > x || x > max-1 {
		return result
	}

	for x != 0 {
		result = result*10 + x%10
		x = x / 10
	}

	if -max >= result || result >= max-1 {
		return 0
	}

	return result
}
