package main

import "fmt"

func main() {
	val := isPalindrome(-121)
	fmt.Printf("%v", val)
}

/*
9. 回文数
给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。例如，121 是回文，而 123 不是。
*/
func isPalindrome(x int) bool {
	// 方法一： 自己写的， 执行消耗大
	/*
		beforeX := x
		var reverse = 0
		max := 1 << 31
		if x < 0 || -max > x || x > max -1 {
			return false
		}

		for x != 0 {
			reverse = reverse *10 + x%10
			x = x/10
		}

		if -max >= reverse || reverse >= max -1 {
			return false
		}

		if beforeX == reverse {
			return true
		}

		return false
	*/

	// 方法二：只反转一般
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}

	// 当数字长度为奇数时，我们可以通过 revertedNumber/10 去除处于中位的数字。
	// 例如，当输入为 12321 时，在 while 循环的末尾我们可以得到 x = 12，revertedNumber = 123，
	// 由于处于中位的数字不影响回文（它总是与自己相等），所以我们可以简单地将其去除。
	return x == revertedNumber || x == revertedNumber/10

}
