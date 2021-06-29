package stringRander

import "testing"

func TestRandString(t *testing.T) {
	type args struct {
		letterBytes string
		n           int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{
				letterBytes: DefaultLetterBytes,
				n:           10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(RandString(tt.args.letterBytes, tt.args.n))
		})
	}
}
