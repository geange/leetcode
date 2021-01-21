package main

import "strings"

func main() {
	convert("PAYPALISHIRING", 4)
}

/**
6. Z 字形变换
将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：

P   A   H   N
A P L S I I G
Y   I   R
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。

请你实现这个将字符串进行指定行数变换的函数：

string convert(string s, int numRows);


示例 1：

输入：s = "PAYPALISHIRING", numRows = 3
输出："PAHNAPLSIIGYIR"
示例 2：
输入：s = "PAYPALISHIRING", numRows = 4
输出："PINALSIGYAHRPI"
解释：
P     I    N
A   L S  I G
Y A   H R
P     I
示例 3：

输入：s = "A", numRows = 1
输出："A"


提示：

1 <= s.length <= 1000
s 由英文字母（小写和大写）、',' 和 '.' 组成
1 <= numRows <= 1000

*/

const (
	VLINE = iota
	DLINE
)

func convert(s string, numRows int) string {
	n := len(s)
	height := numRows
	block := numRows*2 - 2
	width := func() int {
		less := 0
		leftSize := n % block
		if leftSize > numRows {
			less = n%block - numRows + 1
		} else if leftSize > 0 {
			less = 1
		}
		return n/block*(numRows-1) + less
	}()

	cache := make([][]uint8, height)
	for i := 0; i < height; i++ {
		cache[i] = make([]uint8, width)
	}

	x, y := 0, 0
	status := VLINE

	for index, v := range s {

	}

	var buf strings.Builder
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if cache[i][j] != 0 {
				buf.WriteByte(cache[i][j])
			}
		}
	}
	return buf.String()
}

func distance(index, numRows int) (x, y int) {
	//block := numRows*2 - 2
	//x =
	return
}
