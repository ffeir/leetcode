package longest_consecutive_sequence

import "testing"

func TestLongestConsecutiveSequence(t *testing.T) {
	type args struct {
		arrays []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name:"problem_example",
			args:args{
				arrays:[]int{100, 4, 200, 1, 3, 2},
			},
			want:4,
		},
		{
			name:"nil",
			args:args{
				arrays:nil,
			},
			want:0,
		},
		{
			name:"empty",
			args:args{
				arrays:[]int{},
			},
			want:0,
		},
		{
			name:"sorted",
			args:args{
				arrays:[]int{1, 2, 3, 4},
			},
			want:4,
		},
		{
			name:"two_same",
			args:args{
				arrays:[]int{1, 2, 3, 4, 100, 102, 103, 99, 101},
			},
			want:5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestConsecutiveSequence(tt.args.arrays); got != tt.want {
				t.Errorf("LongestConsecutiveSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
