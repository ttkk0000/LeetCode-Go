package leetcode

func twoSum(nums []int, target int) []int {
	//存储上次遍历nums的v，k，如果遇到相同的v则覆盖，可以保证最近的索引输出
	m := make(map[int]int)
	for k, v := range nums {
		if idx, ok := m[target-v]; ok {
			return []int{idx, k}
		}
		m[v] = k
	}
	return nil
}
