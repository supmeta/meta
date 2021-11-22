package main

import "fmt"

func main() {
	//val := isPalindrome(-121)
	//fmt.Printf("%v", val)

	val2 := romanToInt("IV")
	fmt.Printf("%v", val2)
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

/*
13. 罗马数字转整数
罗马数字包含以下七种字符:I，V，X，L，C，D和M。

字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
例如， 罗马数字 2 写做II，即为两个并列的 1 。12 写做XII，即为X+II。 27 写做XXVII, 即为XX+V+II。
通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做IIII，而是IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为IX。这个特殊的规则只适用于以下六种情况：
I可以放在V(5) 和X(10) 的左边，来表示 4 和 9。
X可以放在L(50) 和C(100) 的左边，来表示 40 和90。
C可以放在D(500) 和M(1000) 的左边，来表示400 和900。
给定一个罗马数字，将其转换成整数。
*/
func romanToInt(s string) int {
	result := 0
	covert := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	for i := range s {
		val := covert[string(s[i])]
		if i < len(s)-1 && val < covert[string(s[i+1])] {
			result -= val
		} else {
			result += val
		}
	}
	return result
}
