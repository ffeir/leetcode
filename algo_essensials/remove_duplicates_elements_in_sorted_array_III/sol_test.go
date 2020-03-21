package main

import (
	"testing"
)

func TestRemoveDuplicateArray(t *testing.T) {
	type args struct {
		arrays []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//{
		//	name: "nil",
		//	args: args{
		//		arrays: nil,
		//	},
		//	want: 0,
		//},
		//{
		//	name: "empty",
		//	args: args{
		//		arrays: make([]int, 0),
		//	},
		//	want: 0,
		//},
		//{
		//	name: "only one element",
		//	args: args{
		//		arrays: []int{1},
		//	},
		//	want: 1,
		//},
		//{
		//	name: "two element and same",
		//	args: args{
		//		arrays: []int{1, 1},
		//	},
		//	want: 1,
		//},
		//{
		//	name: "two element and different",
		//	args: args{
		//		arrays: []int{1, 2},
		//	},
		//	want: 2,
		//},
		{
			name: "three element and different",
			args: args{
				arrays: []int{1, 1, 2},
			},
			want: 2,
		},
		{
			name: "three element and different",
			args: args{
				arrays: []int{1, 2, 3},
			},
			want: 3,
		},
		{
			name: "three element",
			args: args{
				arrays: []int{1, 1, 1},
			},
			want: 1,
		},
		{
			name: "three element",
			args: args{
				arrays: []int{1, 1, 1, 2, 2, 2, 3, 3, 3},
			},
			want: 6,
		},
		{
			name: "three element",
			args: args{
				arrays: []int{1, 2, 2, 3, 100, 100},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveDuplicateArray(tt.args.arrays); got != tt.want {
				t.Errorf("RemoveDuplicateArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
