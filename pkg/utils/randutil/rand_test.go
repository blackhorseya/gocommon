package randutil

import (
	"encoding/hex"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	s := String(10)
	assert.Equal(t, 10, len(s))
}

func TestLowerASCII(t *testing.T) {
	s := LowerASCII(10)
	assert.Equal(t, 10, len(s))
}

func TestHex(t *testing.T) {
	s := Hex(10)
	assert.Equal(t, 10, len(s))
	_, e := hex.DecodeString(s)
	assert.NoError(t, e)
}

func TestDigit(t *testing.T) {
	s := Digit(10)
	assert.Equal(t, 10, len(s))
	_, err := strconv.Atoi(s)
	assert.NoError(t, err)
}

func TestStringByCharSet(t *testing.T) {
	s := StringByCharSet(10, "abc")
	assert.Equal(t, 10, len(s))
	c := strings.Count(s, "a") + strings.Count(s, "b") + strings.Count(s, "c")
	assert.Equal(t, 10, c)
}

func BenchmarkString10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		String(10)
	}
}

func BenchmarkString100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		String(100)
	}
}
