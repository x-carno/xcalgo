package strings

import (
	"testing"
)

func TestReverseLeftWords(t *testing.T) {
	type args struct {
		s string
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "normal",
			args: args{
				s: "abcdef",
				n: 2,
			},
			want: "cdefab",
		},
		{
			name: "zero",
			args: args{
				s: "abcdef",
				n: 0,
			},
			want: "abcdef",
		},
		{
			name: "equal",
			args: args{
				s: "abcdef",
				n: 6,
			},
			want: "abcdef",
		},
		{
			name: "bigger",
			args: args{
				s: "abcdef",
				n: 7,
			},
			want: "bcdefa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseLeftWords(tt.args.s, tt.args.n); got != tt.want {
				t.Errorf("ReverseLeftWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFirstUniqChar(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		// TODO: Add test cases.
		{
			name: "no",
			args: args{
				s: "leetcodea",
			},
			want: 'l',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstUniqChar(tt.args.s); got != tt.want {
				t.Errorf("FirstUniqChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "y",
			args: args{
				haystack: "aaab",
				needle:   "aab",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("StrStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				num1: "30",
				num2: "1",
			},
			want: "30",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Multiply(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddBinary(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				a: "10",
				b: "100",
			},
			want: "110",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddBinary(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("AddBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLengthOfLastWord(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{s: "Hello World"},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthOfLastWord(tt.args.s); got != tt.want {
				t.Errorf("LengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
