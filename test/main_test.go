package main

import (
	"reflect"
	"testing"
)

func TestEvalRPN(t *testing.T) {
	type args struct {
		tokens []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "tt",
			args: args{
				tokens: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			},
			want: 22,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EvalRPN(tt.args.tokens); got != tt.want {
				t.Errorf("EvalRPN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlusOne(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{
				digits: []int{0},
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PlusOne(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{n: 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fib(tt.args.n); got != tt.want {
				t.Errorf("Fib() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDailyTemperatures(t *testing.T) {
	type args struct {
		temperatures []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{temperatures: []int{90, 60, 90}},
			want: []int{0, 1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DailyTemperatures(tt.args.temperatures); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DailyTemperatures() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasAlternatingBits(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		// {
		// 	name: "t1",
		// 	args: args{n: 1366},
		// 	want: false,
		// },
		// {
		// 	name: "t1",
		// 	args: args{n: 1},
		// 	want: true,
		// },
		{
			name: "t1",
			args: args{n: 1365},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasAlternatingBits(tt.args.n); got != tt.want {
				t.Errorf("HasAlternatingBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
