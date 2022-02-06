package miscellaneous

import (
	"testing"
)

func TestSievePrime(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{
				n: 0,
			},
		},
		{
			name: "t2",
			args: args{
				n: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SievePrimeConcurrent(tt.args.n)
		})
	}
}

func TestSievePrimeNormal(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{
				n: 0,
			},
		},
		{
			name: "t1",
			args: args{
				n: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SievePrimeNormal(tt.args.n)
		})
	}
}
