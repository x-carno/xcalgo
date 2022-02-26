package arrays

import (
	"math/rand"
	"testing"
	"time"
)

const BigCount = 10

var testArr []int
var testMissingArr []int

func init() {
	for i := 0; i < BigCount; i++ {
		testArr = append(testArr, i+1)
	}

	rand.Seed(time.Now().Unix())
	miss := rand.Intn(BigCount)
	for i := 0; i < miss; i++ {
		testMissingArr = append(testArr, i)
	}
	for i := miss + 1; i < BigCount; i++ {
		testMissingArr = append(testArr, i)
	}
}

func BenchmarkSearchRepeatCount(b *testing.B) {
	// arr := []int{2, 3, 3, 5, 7, 8, 8, 8, 10}
	for i := 0; i < b.N; i++ {
		SearchRepeatCount(testArr, 57)
	}
}

func BenchmarkSearchRepeatCountNormal(b *testing.B) {
	// arr := []int{2, 3, 3, 5, 7, 8, 8, 8, 10}
	for i := 0; i < b.N; i++ {
		SearchRepeatCountNormal(testArr, 57)
	}
}

func BenchmarkMissingNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MissingNumber(testMissingArr)
	}
}
