package trees

import (
	"reflect"
	"testing"
)

func TestComputeTreeLayer(t *testing.T) {
	type args struct {
		arr []interface{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "nil",
			args: args{
				arr: nil,
			},
			want: 0,
		},
		{
			name: "1",
			args: args{
				arr: []interface{}{1},
			},
			want: 1,
		},
		{
			name: "3",
			args: args{
				arr: []interface{}{1, 2, 3, 4},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := computeTreeLayer(tt.args.arr); got != tt.want {
				t.Errorf("ComputeTreeLayer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTreeNode_LevelOrder(t *testing.T) {
	type fields struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		// TODO: Add test cases.
		{
			name: "1",
			fields: fields{
				Val: 1,
				Left: &TreeNode{
					Val:   3,
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{
					Val:   5,
					Right: &TreeNode{Val: 6},
				},
			},
			want: []int{1, 3, 5, 4, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &TreeNode{
				Val:   tt.fields.Val,
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			if got := r.LevelOrder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TreeNode.LevelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
