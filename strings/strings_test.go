package strings

import "testing"

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
