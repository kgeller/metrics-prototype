package main

import (
	"testing"

	"gotest.tools/assert"
)

func TestGetMetricValue(t *testing.T) {
	var testMetric = map[string]interface{}{
		"histogramMetric": map[string]map[string]float64{
			"histogram": {
				"metric": float64(5),
			},
		},
		"numerical": float64(10),
	}

	assert.Equal(t, getMetricValue("histogramMetric.histogram.metric", testMetric), float64(5))
	assert.Equal(t, getMetricValue("numerical", testMetric), float64(10))
}
