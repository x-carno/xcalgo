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

func TestTreeNode_LevelOrder2D(t *testing.T) {
	type fields struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}
	tests := []struct {
		name   string
		fields fields
		want   [][]int
	}{
		// TODO: Add test cases.
		{
			name: "pass",
			fields: fields{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			want: [][]int{{3}, {9, 20}, {15, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := &TreeNode{
				Val:   tt.fields.Val,
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			if got := root.LevelOrder2D(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TreeNode.LevelOrder2D() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTreeNode_LevelOrder2DReverseOdd(t *testing.T) {
	type fields struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}
	tests := []struct {
		name   string
		fields fields
		want   [][]int
	}{
		// TODO: Add test cases.
		{
			name: "pass",
			fields: fields{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			want: [][]int{{3}, {20, 9}, {15, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := &TreeNode{
				Val:   tt.fields.Val,
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			if got := root.LevelOrder2DReverseOdd(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TreeNode.LevelOrder2DReverseOdd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTreeNode_IsSymmetric(t *testing.T) {
	type fields struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
		{
			name: "pass",
			fields: fields{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
					Left: &TreeNode{
						Val: 7,
					},
					Right: &TreeNode{
						Val: 15,
					},
				},
				Right: &TreeNode{
					Val: 9,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := &TreeNode{
				Val:   tt.fields.Val,
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			if got := root.IsSymmetric(); got != tt.want {
				t.Errorf("TreeNode.IsSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTreeNode_IsSubStructure(t *testing.T) {
	type fields struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}
	type args struct {
		B *TreeNode
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{
			name: "pass",
			fields: fields{
				Val: 4,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   4,
						Left:  &TreeNode{Val: 8},
						Right: &TreeNode{Val: 9},
					},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 7},
				},
			},
			args: args{
				B: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 8},
					Right: &TreeNode{Val: 9},
				},
			},
			want: true,
		},
		{
			name: "pass",
			fields: fields{
				Val: 1,
				Left: &TreeNode{
					Val:   0,
					Left:  &TreeNode{Val: -4},
					Right: &TreeNode{Val: -3},
				},
				Right: &TreeNode{Val: 1},
			},
			args: args{
				B: &TreeNode{
					Val:  1,
					Left: &TreeNode{Val: -4},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			A := &TreeNode{
				Val:   tt.fields.Val,
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			if got := A.IsSubStructure(tt.args.B); got != tt.want {
				t.Errorf("TreeNode.IsSubStructure() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTreeNode_isPrefixSequence(t *testing.T) {
	type fields struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}
	type args struct {
		B *TreeNode
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{
			name: "pass",
			fields: fields{
				Val: 4,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   4,
						Left:  &TreeNode{Val: 8},
						Right: &TreeNode{Val: 9},
					},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 7},
				},
			},
			args: args{
				B: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 8},
					Right: &TreeNode{Val: 9},
				},
			},
			want: false,
		},
		{
			name: "pass",
			fields: fields{
				Val: 4,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   4,
						Left:  &TreeNode{Val: 8},
						Right: &TreeNode{Val: 9},
					},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 7},
				},
			},
			args: args{
				B: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 2},
					Right: &TreeNode{Val: 3},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			A := &TreeNode{
				Val:   tt.fields.Val,
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			if got := A.hasPrefixSequence(tt.args.B); got != tt.want {
				t.Errorf("TreeNode.isPrefixSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTreeNode_Height(t *testing.T) {
	type fields struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			fields: fields{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val: 20,
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := &TreeNode{
				Val:   tt.fields.Val,
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			if got := root.Height(); got != tt.want {
				t.Errorf("TreeNode.Height() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTreeNode_IsBalanced(t *testing.T) {
	type fields struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			fields: fields{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := &TreeNode{
				Val:   tt.fields.Val,
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			if got := root.IsBalanced(); got != tt.want {
				t.Errorf("TreeNode.IsBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
