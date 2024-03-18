package problem2

import "testing"

func TestSupermanChickenRescue(t *testing.T) {
	type args struct {
		chickenCount int
		lengthOfRoof int
		chickenPos   []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Case 1: 5 chickens, 5 length of roof, [2 5 10 12 15]",
			args: args{
				chickenCount: 5,
				lengthOfRoof: 5,
				chickenPos:   []int{2, 5, 10, 12, 15},
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "Case 2: 6 chickens, 10 length of roof, [1 11 30 34 35 37]",
			args: args{
				chickenCount: 6,
				lengthOfRoof: 10,
				chickenPos:   []int{1, 11, 30, 34, 35, 37},
			},
			want:    4,
			wantErr: false,
		},
		{
			name: "Case 3: 6 chickens, 10 length of roof, [1 11 30]",
			args: args{
				chickenCount: 6,
				lengthOfRoof: 10,
				chickenPos:   []int{1, 11, 30},
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SupermanChickenRescue(tt.args.chickenCount, tt.args.lengthOfRoof, tt.args.chickenPos)
			if (err != nil) != tt.wantErr {
				t.Errorf("SupermanChickenRescue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SupermanChickenRescue() got = %v, want %v", got, tt.want)
			}
		})
	}
}
