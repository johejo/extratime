package extratime

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_cut(t *testing.T) {
	tests := []struct {
		in, out string
	}{
		{`"hello"`, `hello`},
		{`world`, `world`},
	}
	for _, tt := range tests {
		t.Run(tt.in+`->`+tt.out, func(t *testing.T) {
			got := cut([]byte(tt.in))
			assert.Equal(t, tt.out, got)
		})
	}
}
