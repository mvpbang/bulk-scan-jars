package main

func maxArea(height []int) int {
	res := 0
	//max := []int
	max := make([]int, 0)
	// 迭代int slice必须make初始化
	// error:runtime error:index out of range [0] with length 0
	for i, _ := range max {
		if res < max[i] {
			res = max[i]
		}
	}
	return res
}

func main() {
	var height []int
	maxArea(height)

}
