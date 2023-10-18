package main

type operator string

const (
	GreaterThan       operator = "gt"
	GreaterThanEquals operator = "gte"
	LessThan          operator = "lt"
	LessThanEquals    operator = "lte"
)

type Condition struct {
	MetricName string
	Operator   operator
	Threshold  float64
}

type Check struct {
	Input      string
	DocLink    string
	Conditions []Condition
}

type MetricWarning struct {
	InputName string
	Check     Check
}

type HistogramMetric map[string]map[string]float64
