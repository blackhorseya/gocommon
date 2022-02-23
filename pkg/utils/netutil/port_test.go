package netutil

import "testing"

func TestGetAvailablePort(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "get port then success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAvailablePort(); !(0 <= got && got <= 65535) {
				t.Errorf("GetAvailablePort() = %v, the got port out of range 0-65535", got)
			}
		})
	}
}
