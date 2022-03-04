package strings

import (
	"fmt"
	"strings"
)

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

// num1 and num2 are integer strings
// e.g. num1="123" num2="234"
// return "28782"
func Multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	if len(num2) < len(num1) {
		num1, num2 = num2, num1
	}
	if num1 == "1" {
		return num2
	}
	l1, l2 := len(num1), len(num2)
	ret := make([]int, l1+l2)
	for i := l1 - 1; i >= 0; i-- {
		for j := l2 - 1; j >= 0; j-- {
			retIdx := (l1 - 1 - i) + (l2 - 1 - j)
			c := int(num1[i]-'0') * int(num2[j]-'0')
			fmt.Println(c)
			for {
				c = c + ret[retIdx]
				ret[retIdx] = c % 10
				c = c / 10
				retIdx++
				if c == 0 {
					break
				}
			}
		}
	}
	if ret[len(ret)-1] == 0 {
		ret = ret[:len(ret)-1]
	}
	retS := strings.Builder{}
	for i := len(ret) - 1; i >= 0; i-- {
		retS.WriteByte(byte(ret[i] + '0'))
	}
	return retS.String()
}
