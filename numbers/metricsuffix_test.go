package numbers

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestMetricSuffix(t *testing.T) {
    tests := map[float64]string{
        -1: "-1",
        -1000: "-1k",
        0: "0",
        1: "1",
        101: "101",
        1000: "1k",
        1240: "1.2k",
        1240000: "1.24M",
        3500000: "3.5M",
    }
    
    for input, expected := range tests {
        assert.Equal(t, expected, MetricSuffix(input))
    }
}
