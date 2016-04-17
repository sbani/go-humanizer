package units

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySuffix(t *testing.T) {
	tests := map[float64]string{
		0:                 "0 bytes",
		1:                 "1 bytes",
		1024:              "1.0 kB",
		1025:              "1.0 kB",
		1536:              "1.5 kB",
		1048576 * 5:       "5.00 MB",
		1073741824 * 2:    "2.00 GB",
		1099511627776 * 3: "3.00 TB",
		1325899906842624:  "1.18 PB",
		-5:                "-5 bytes",
		-1.49:             "-1 bytes",
		-2.8:              "-3 bytes",
	}

	for input, expected := range tests {
		assert.Equal(t, expected, BinarySuffix(input))
	}
}
