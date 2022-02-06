package sort

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func BenchmarkTestMergeSort_Sort(b *testing.B) {
	// s := []int{1, 3, 4, 5, 7, 9, 3}
	for i := 0; i < b.N; i++ {
		ms := NewMergeSort(testSlice)
		ms.Sort()
	}
}

var originTestSlice []int
var testSlice []int

func init() {
	for i := 0; i < 100000; i++ {
		originTestSlice = append(originTestSlice, i+1)
		testSlice = append(testSlice, i+1)
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(testSlice), func(i, j int) { testSlice[i], testSlice[j] = testSlice[j], testSlice[i] })
}

func TestMergeSort_Sort(t *testing.T) {
	type fields struct {
		items []int
		// size       int
		// left       int
		// mid        int
		// right      int
		// innerSlice []int
		// grPool     *ants.Pool
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{
			name: "one number",
			fields: fields{
				items: []int{2},
			},
			want: []int{2},
		},
		{
			name: "two number",
			fields: fields{
				items: []int{2, 1},
			},
			want: []int{1, 2},
		},
		{
			name: "three number",
			fields: fields{
				items: []int{3, 1, 2},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "sorted",
			fields: fields{
				items: []int{1, 2, 3, 4, 5, 6},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "reversed",
			fields: fields{
				items: []int{5, 4, 3, 2, 1},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "randomed",
			fields: fields{
				items: []int{4, 2, 3, 5, 1},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "sorted",
			fields: fields{
				items: []int{8, 2, 4, 6, 9, 1, 7, 3, 5, 10},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name: "long size",
			fields: fields{
				items: testSlice,
			},
			want: originTestSlice,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMergeSort(tt.fields.items)
			if got := m.Sort(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort.Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeSort_sort3(t *testing.T) {
	type fields struct {
		items []int
	}
	type args struct {
		le int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name: "t1",
			fields: fields{
				items: []int{1, 2, 3, 4},
			},
			args: args{
				le: 1,
			},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMergeSort(tt.fields.items)
			m.sort3(tt.args.le)
			if !reflect.DeepEqual(m.items, tt.want) {
				t.Errorf("MergeSort.sort3() = %v, want %v", m.items, tt.want)
			}
		})
	}
}
