package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/x-carno/xcalgo/dynamic"
	"github.com/x-carno/xcalgo/util"
)

var (
	f = func(x int) int {
		for i := 0; i < 1_0000_0000; i++ {
			x++
		}
		return x
	}
)

func main() {

	fmt.Println(1 << 62)
	util.TimeFunc(func() {
		fmt.Println(dynamic.TranslateNum(1111111111111111111))
	})

	matrix := [][]int{
		{1, 2, 3, 4, 5},
		{2, 3, 4, 5, 6},
		{3, 4, 5, 6, 7},
		{4, 5, 6, 7, 8},
		{5, 6, 7, 8, 9},
	}
	fmt.Println(matrix)

	// xx:=matrix[0:2]
	sub1 := matrix[0][3:5]
	sub2 := matrix[1][3:5]
	fmt.Println(sub1)
	fmt.Println(sub2)
	subM := [][]int{}
	subM = append(subM, sub1)
	subM = append(subM, sub2)
	fmt.Println(subM)

	var b strings.Builder
	str := "sj dfks jdf"
	for _, v := range str {
		fmt.Println(v)
		if v == 32 {
			b.WriteString("%20")
		} else {
			b.WriteRune(v)
		}
	}
	fmt.Println(b.String())

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)
	var inputs []string

	var length int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &length)

	scanner.Scan()
	inputs = strings.Split(scanner.Text(), " ")
	for i := 0; i < length; i++ {
		part1 := inputs[i]
		fmt.Println(part1)
	}
	inp1 := inputs
	fmt.Println(inputs)
	scanner.Scan()
	inputs = strings.Split(scanner.Text(), " ")
	for i := 0; i < length; i++ {
		part2 := inputs[i]
		fmt.Println(part2)
	}
	inp2 := inputs
	fmt.Println(inputs)
	AA(inp1, inp2)
}

func AA(p1, p2 []string) {
	for i := 0; i < len(p1); i++ {
		noneE := false
		for j := 0; j < len(p2); j++ {
			if p1[i] == p2[j] {
				fmt.Println(j)
				noneE = true
				break
			}
		}
		if !noneE {
			fmt.Println("NONE")
		}
	}
}

// 逆波兰表达式求值
func EvalRPN(tokens []string) int {
	ret := make([]int, 0)
	for _, v := range tokens {
		switch v {
		case "+":
			m := ret[len(ret)-2] + ret[len(ret)-1]
			ret = append(ret[:len(ret)-2], m)
		case "-":
			m := ret[len(ret)-2] - ret[len(ret)-1]
			ret = append(ret[:len(ret)-2], m)
		case "*":
			m := ret[len(ret)-2] * ret[len(ret)-1]
			ret = append(ret[:len(ret)-2], m)
		case "/":
			m := ret[len(ret)-2] / ret[len(ret)-1]
			ret = append(ret[:len(ret)-2], m)
		default:
			n, _ := strconv.Atoi(v)
			ret = append(ret, n)
		}
	}
	return ret[0]
}

func PlusOne(digits []int) []int {
	digits = append(digits, 0)
	copy(digits[1:], digits)
	digits[0] = 0
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i]++
			break
		}
	}
	if digits[0] == 0 {
		digits = digits[1:]
	}
	return digits
}

func Fib(n int) int {
	if n == 0 {
		return 0
	}
	cm := []int{0, 1}
	for i := 0; i < n-1; i++ {
		cm[0] = cm[0] + cm[1]
		cm[0], cm[1] = cm[1], cm[0]
	}
	return cm[1]
}

// 给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，
// 其中 answer[i] 是指在第 i 天之后，才会有更高的温度。如果气温在这之后都不会升高，请在该位置用 0 来代替。
// 输入: temperatures = [73,74,75,71,69,72,76,73]
// 输出: [1,1,4,2,1,1,0,0]
func DailyTemperatures(temperatures []int) []int {
	ret := make([]int, len(temperatures))
	for i := len(temperatures) - 2; i >= 0; i-- {
		if temperatures[i] < temperatures[i+1] {
			ret[i] = 1
			continue
		}
		j := i + 1 + ret[i+1]
		for j < len(temperatures) {
			if temperatures[i] < temperatures[j] {
				ret[i] = j - i
				break
			}
			if ret[j] == 0 {
				break
			}
			j += ret[j]
		}
	}
	return ret
}

func Rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			matrix[i][j], matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1] =
				matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1], matrix[i][j]
		}
	}
}

func FindRotation(mat [][]int, target [][]int) bool {
	compare := func(matInner [][]int, targetInner [][]int) bool {
		for i := 0; i < len(targetInner); i++ {
			for j := 0; j < len(matInner); j++ {
				if matInner[i][j] != targetInner[i][j] {
					return false
				}
			}
		}
		return true
	}
	n := len(mat)
	for r := 0; r < 4; r++ {
		if compare(mat, target) {
			return true
		}
		for i := 0; i < n/2; i++ {
			for j := 0; j < (n+1)/2; j++ {
				mat[i][j], mat[n-j-1][i], mat[n-i-1][n-j-1], mat[j][n-i-1] =
					mat[n-j-1][i], mat[n-i-1][n-j-1], mat[j][n-i-1], mat[i][j]
			}
		}
	}
	return false
}

func HasAlternatingBits(n int) bool {
	last := n % 2
	for n > 0 {
		n = n >> 1
		pre := n % 2
		if pre == last {
			return false
		}
		last = pre
	}
	return true
}
