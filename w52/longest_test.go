package w52

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	s := "international"
	len := LengthOfLongestSubstring(s)

	if len != 7 {
		t.Errorf("expected: 7('ernatio') from '%s', but got: %d", s, len)
	}
}
