package strings

// s = "abcdefg", n = 2
// return cdefgab
func ReverseLeftWords(s string, n int) string {
	k := n % len(s)
	i := s[:k]
	j := s[k:]
	return j + i
}
