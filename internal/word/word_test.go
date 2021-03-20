/*
@Time : 2021/3/20 10:59
@Author : Tux
@Description :
*/

package word

import (
	"testing"
)

func TestUnderscoreToUpperCamelCase(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "no1",
			args: args{s: "a_b_c"},
			want: "ABC",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnderscoreToUpperCamelCase(tt.args.s); got != tt.want {
				t.Errorf("UnderscoreToUpperCamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
	t.Log(UnderscoreToLowerCamelCase("a_b_c"))
	t.Log(CamelCaseToUnderscore("AB&*C"))

}
