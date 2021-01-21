package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestPalindrome("cbbd"))
}

/**
5. 最长回文子串
给你一个字符串 s，找到 s 中最长的回文子串。



示例 1：

输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
示例 2：

输入：s = "cbbd"
输出："bb"
示例 3：

输入：s = "a"
输出："a"
示例 4：

输入：s = "ac"
输出："a"


提示：

1 <= s.length <= 1000
s 仅由数字和英文字母（大写和/或小写）组成

*/

func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]byte, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]byte, len(s))
	}

	ans := ""

	for l := 0; l < n; l++ {
		for i := 0; i+l < n; i++ {
			j := i + l
			switch l {
			case 0:
				dp[i][j] = 1
			case 1:
				if s[i] == s[j] {
					dp[i][j] = 1
				}
			default:
				if s[i] == s[j] {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			if dp[i][j] > 0 && l+1 > len(ans) {
				ans = s[i : i+l+1]
			}
		}
	}

	return ans
}
