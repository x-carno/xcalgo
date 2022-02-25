package arrays

import (
	"testing"
)

func TestFindRepeatNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "yes",
			args: args{
				nums: []int{3, 2, 1, 2, 3},
			},
			want: 2,
		},
		{
			name: "no repeat",
			args: args{
				nums: []int{3, 2, 1},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindRepeatNumber(tt.args.nums); got != tt.want {
				t.Errorf("FindRepeatNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchRepeatCount(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "pass1",
			args: args{
				nums:   []int{2, 3, 3, 5, 7, 8, 8, 8, 10},
				target: 2,
			},
			want: 1,
		},
		{
			name: "pass2",
			args: args{
				nums:   []int{2, 3, 3, 5, 7, 8, 8, 8, 10},
				target: 3,
			},
			want: 2,
		},
		{
			name: "pass3",
			args: args{
				nums:   []int{2, 3, 3, 5, 7, 8, 8, 8, 10},
				target: 8,
			},
			want: 3,
		},
		{
			name: "not in array",
			args: args{
				nums:   []int{2, 3, 3, 5, 7, 8, 8, 8, 10},
				target: 4,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchRepeatCount(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("SearchRepeatCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMissingNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "one",
			args: args{
				nums: []int{0},
			},
			want: 1,
		},
		{
			name: "two",
			args: args{
				nums: []int{0, 1},
			},
			want: 2,
		},
		{
			name: "three",
			args: args{
				nums: []int{0, 1, 2, 3, 5, 6, 7, 8, 9, 10},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MissingNumber(tt.args.nums); got != tt.want {
				t.Errorf("MissingNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
