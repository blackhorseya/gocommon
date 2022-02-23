package stringutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountIntersectStringSlice(t *testing.T) {
	type args struct {
		s1 []string
		s2 []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[a, b, c] [b, c] then 2",
			args: args{s1: []string{"a", "b", "c"}, s2: []string{"b", "c"}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, CountIntersectStringSlice(tt.args.s1, tt.args.s2), "CountIntersectStringSlice(%v, %v)", tt.args.s1, tt.args.s2)
		})
	}
}

func TestInterfaceSliceToStringSlice(t *testing.T) {
	type args struct {
		params []interface{}
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "[a, b] then [a, b]",
			args: args{params: []interface{}{"a", "b"}},
			want: []string{"a", "b"},
		},
		{
			name: "[a, 1, b] then [a, b]",
			args: args{params: []interface{}{"a", 1, "b"}},
			want: []string{"a", "b"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, InterfaceSliceToStringSlice(tt.args.params...), "InterfaceSliceToStringSlice(%v)", tt.args.params...)
		})
	}
}
