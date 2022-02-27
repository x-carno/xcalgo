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

func StrStr(haystack string, needle string) int {
	// strings.Index(haystack, needle)
	if len(needle) == 0 {
		return 0
	}
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		flag := true
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				flag = false
				break
			}
		}

		if flag {
			return i
		}
	}
	return -1
}
