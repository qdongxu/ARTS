package w52

// LengthOfLongestSubstring is a solution for https://leetcode.com/problems/longest-substring-without-repeating-characters/
func LengthOfLongestSubstring(s string) int {
	size := len(s)
	chars := make(map[byte]bool)
	i, j, max := 0, 0, 0
	for i < size && j < size {
		_, ok := chars[s[j]]
		if ! ok {
			chars[s[j]] = true
			j++

			if max < j-i {
				max = j - i
			}
		} else {
			delete(chars, s[i])
			i++

		}
	}

	return max
}
