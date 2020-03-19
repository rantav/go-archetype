package reader

import "testing"

func Test_relative(t *testing.T) {
	type args struct {
		prefix string
		path   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty prefix",
			args: args{
				prefix: "",
				path:   "a",
			},
			want: "a",
		},
		{
			name: "prefix .",
			args: args{
				prefix: ".",
				path:   "a",
			},
			want: "a",
		},
		{
			name: "prefix ../../b/",
			args: args{
				prefix: "../../b/",
				path:   "../../b/a",
			},
			want: "a",
		},
		{
			name: "prefix ../../b",
			args: args{
				prefix: "../../b",
				path:   "../../b/a",
			},
			want: "a",
		},
	}
	for _, td := range tests {
		tt := td
		t.Run(tt.name, func(t *testing.T) {
			if got := relative(tt.args.prefix, tt.args.path); got != tt.want {
				t.Errorf("relative() = %v, want %v", got, tt.want)
			}
		})
	}
}
