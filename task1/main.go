package main

import (
	"sort"
	"strconv"
)

func singleNumber1(nums []int) int {
	var res int
	for _, v := range nums {
		res ^= v
	}
	return res
}

func singleNumber2(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return -1
}

/**
 * 判断是否是回文数
 */
func isPalindrome(x int) bool {

	//将x转为字符串
	s := strconv.Itoa(x)
	//翻转字符串
	s2 := reverseString(s)
	return s == s2

}

func reverseString(s string) interface{} {
	var res string
	for i := len(s) - 1; i >= 0; i-- {
		res += string(s[i])
	}
	return res
}

/*
*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。
*/
func isValid(s string) bool {
	// 定义一个栈
	stack := make([]byte, 0)

	// 定义括号匹配映射
	mapping := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	for i := 0; i < len(s); i++ {
		char := s[i]

		// 如果是右括号
		if _, exists := mapping[char]; exists {
			// 检查栈是否为空或者栈顶是否匹配
			if len(stack) == 0 || stack[len(stack)-1] != mapping[char] {
				return false
			}
			// 弹出栈顶
			stack = stack[:len(stack)-1]
		} else {
			// 如果是左括号，压入栈
			stack = append(stack, char)
		}
	}

	// 栈为空说明所有括号都匹配
	return len(stack) == 0
}

/*
*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。
*/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	// 公共前缀比所有字符串都短，随便选一个先
	s := strs[0]
	for _, str := range strs {
		for !startsWith(str, s) {
			if len(s) == 0 {
				return ""
			}
			// 公共前缀不匹配就让它变短！
			s = s[:len(s)-1]
		}
	}
	return s
}

// 辅助函数：判断字符串 s 是否以 prefix 开头
func startsWith(s string, prefix string) bool {
	if len(s) < len(prefix) {
		return false
	}
	return s[:len(prefix)] == prefix
}

/*
*
给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。

将大整数加 1，并返回结果的数字数组。
*/
func plusOne(digits []int) []int {
	//假如数组是[]int{1, 2, 3})，把数组转成数字123，然后+1，变成124，再转成数组[]int{1, 2, 4})

	num := 0
	for i := 0; i < len(digits); i++ {
		num = num*10 + digits[i]
	}
	num++
	//数字123转成数组[]int{1, 2, 3}) 放到一个新数组
	digits = []int{}
	for num > 0 {
		digits = append([]int{num % 10}, digits...)
		num = num / 10
	}

	return digits

}

/**
给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。
*/

func removeDuplicates(nums []int) int {
	//1. 创建一个新数组，把nums数组的元素放到新数组中，新数组的元素是唯一的
	var newNums []int
	for i := 0; i < len(nums); i++ {
		if !contains(newNums, nums[i]) {
			newNums = append(newNums, nums[i])
		}
	}
	return len(newNums)

}

func contains(nums []int, i int) bool {
	for _, v := range nums {
		if v == i {
			return true
		}
	}
	return false
}

/*
*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。intervals = [[1,3],[2,6],[8,10],[15,18]]
*/
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	// 按起始点排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var newNums [][]int
	newNums = append(newNums, intervals[0])

	for i := 1; i < len(intervals); i++ {
		last := newNums[len(newNums)-1]
		currentStart := intervals[i][0]
		currentEnd := intervals[i][1]

		// 如果当前区间与上一个区间重叠，则合并
		if currentStart <= last[1] {
			newNums[len(newNums)-1] = []int{last[0], max(last[1], currentEnd)}
		} else {
			newNums = append(newNums, []int{currentStart, currentEnd})
		}
	}

	return newNums
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
*
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
*/
func twoSum(nums []int, target int) []int {

	index := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				// 返回索引
				index = append(index, i, j)
			}
		}
	}
	return index

}

func main() {
	println(singleNumber1([]int{2, 2, 1}))
	println(singleNumber2([]int{4, 1, 2, 1, 2}))
	println(isPalindrome(123))
	println(isValid("([])"))
	println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	println(plusOne([]int{1, 2, 3}))
	println(removeDuplicates([]int{1, 1, 2}))
	println(merge([][]int{[]int{1, 3}, []int{2, 6}, []int{8, 10}, []int{15, 18}}))
	println(twoSum([]int{5, 2, 2, 3}, 8))

}
