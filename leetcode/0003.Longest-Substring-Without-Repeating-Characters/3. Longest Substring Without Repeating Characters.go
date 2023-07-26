package leetcode

// 解法一 位图
func lengthOfLongestSubstring(s string) (int, string) {
	if len(s) == 0 {
		return 0, ""
	}
	var bitSet [256]bool
	result, left, right, ss := 0, 0, 0, ""
	for left < len(s) {
		// 右侧字符对应的 bitSet 被标记 true，说明此字符在 X 位置重复，需要左侧向前移动，直到将 X 标记为 false
		var hasSame bool
		if bitSet[s[right]] {
			bitSet[s[left]] = false
			left++
			hasSame = false
		} else {
			bitSet[s[right]] = true
			right++
			hasSame = true
		}
		if hasSame && result < right-left {
			result = right - left
			ss = s[left:right]
		}
		if left+result >= len(s) || right >= len(s) {
			break
		}
	}
	return result, ss
}

// 解法二 滑动窗口
func lengthOfLongestSubstring1(s string) (int, string) {
	if len(s) == 0 {
		return 0, ""
	}
	var freq [127]int
	result, left, right, ss := 0, 0, -1, ""

	for left < len(s) {
		if right+1 < len(s) && freq[s[right+1]] == 0 {
			freq[s[right+1]]++
			right++
		} else {
			freq[s[left]]--
			left++
		}
		old := result
		result = max(result, right-left+1)
		if result > old {
			ss = s[left : right+1]
		}
	}
	return result, ss
}

// 解法三 滑动窗口-哈希桶
func lengthOfLongestSubstring2(s string) (int, string) {
	left, right, res, ss := 0, 0, 0, ""
	indexes := make(map[byte]int, len(s))
	for right < len(s) {
		if idx, ok := indexes[s[right]]; ok && idx >= left {
			left = idx + 1
		}
		indexes[s[right]] = right
		right++
		old := res
		res = max(res, right-left)
		if res > old {
			ss = s[left:right]
		}
	}
	return res, ss
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
