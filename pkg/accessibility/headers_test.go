package accessibility

import "testing"

func TestHeaders_isIncorrectLevel(t *testing.T) {
	type args struct {
		last   int
		actual int
	}

	tests := []struct {
		name string
		h    *Headers
		args args
		want bool
	}{

		{
			name: "Invalid header with 1, 3",
			args: args{1, 3},
			want: true,
		},

		{
			name: "Invalid header with 2, 6",
			args: args{2, 6},
			want: true,
		},

		{
			name: "Correct level with 1, 2",
			args: args{1, 2},
			want: false,
		},

		{
			name: "Valid header level with 1, 1",
			args: args{1, 1},
			want: false,
		},

		{
			name: "valid header with 2, 3",
			args: args{2, 3},
			want: false,
		},

		{
			name: "Valid text with 3, 2",
			args: args{3, 2},
			want: false,
		},

		{
			name: "Valide headers with 4, 1",
			args: args{4, 1},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.isIncorrectLevel(tt.args.last, tt.args.actual); got != tt.want {
				t.Errorf("Headers.isIncorrectLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}
