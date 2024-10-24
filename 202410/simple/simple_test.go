package simple

import "testing"

func Test_add(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				i: 1,
				j: 2,
			},
			want: 3,
		}, {
			name: "2",
			args: args{
				i: 2,
				j: 3,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}
