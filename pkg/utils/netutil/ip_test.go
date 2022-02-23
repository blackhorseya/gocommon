package netutil

import (
	"regexp"
	"testing"
)

var (
	compile, _ = regexp.Compile(`(\b25[0-5]|\b2[0-4][0-9]|\b[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)
)

func TestGetLocalIP4(t *testing.T) {

	tests := []struct {
		name string
	}{
		{
			name: "get local ipv4 then success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIP := GetLocalIP4(); !compile.MatchString(gotIP) {
				t.Errorf("GetLocalIP4() = %v is not ip", gotIP)
			}
		})
	}
}
