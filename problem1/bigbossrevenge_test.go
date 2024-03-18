package problem1

import (
	"testing"
)

func TestBossBabyRevenge(t *testing.T) {
	//oneMInput := "S"
	//for i := 1; i < 1000000; i++ {
	//	oneMInput += "R"
	//	if i%1000 == 0 {
	//		fmt.Println(i)
	//	}
	//}
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Case 1: SRSSRRR",
			args: args{"SRSSRRR"},
			want: "Good Boy",
		},
		{
			name: "Case 2: RSSRR",
			args: args{"RSSRR"},
			want: "Bad Boy",
		},
		{
			name: "Case 3: SSSRRRRS",
			args: args{"SSSRRRRS"},
			want: "Bad Boy",
		},
		{
			name: "Case 4: SSRR",
			args: args{"SSRR"},
			want: "Good Boy",
		},
		{
			name: "Case 5: SRRSSR",
			args: args{"SRRSSR"},
			want: "Bad Boy",
		},
		//{
		//	name: "Case 6: 1M R",
		//	args: args{oneMInput},
		//	want: "Good Boy",
		//},
		{
			name: "Case Invalid",
			args: args{"SSSSSSSSInvalid"},
			want: "Invalid input",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BossBabyRevenge(tt.args.input); got != tt.want {
				t.Errorf("BossBabyRevenge() = %v, want %v", got, tt.want)
			}
		})
	}
}
