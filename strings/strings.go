package strings

// s = "abcdefg", n = 2
// return cdefgab
func ReverseLeftWords(s string, n int) string {
	k := n % len(s)
	i := s[:k]
	j := s[k:]
	return j + i
}

// e.g. "abaccdeff"
// return b
func FirstUniqChar(s string) byte {
	cnt := [26]int{}
	for _, ch := range s {
		cnt[ch-'a']++
	}
	for i, ch := range s {
		if cnt[ch-'a'] == 1 {
			return s[i]
		}
	}
	return ' '
}
