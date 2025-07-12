package main

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

func main() {
	println(singleNumber1([]int{2, 2, 1}))
	println(singleNumber2([]int{4, 1, 2, 1, 2}))
}
