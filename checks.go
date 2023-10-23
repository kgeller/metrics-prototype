package main

var checks = []Check{
	{
		Input: "aws-s3",
		Conditions: []Condition{
			{
				MetricName: "sqs_worker_utilization",
				Operator:   GreaterThan,
				Threshold:  0.9,
			},
			{
				MetricName: "sqs_lag_time.histogram.median",
				Operator:   GreaterThan,
				Threshold:  300000000000, // 5m -> ns
			},
		},
		DocLink: "./guides/001",
	},
	{
		Input: "udp",
		Conditions: []Condition{
			{
				MetricName: "received_bytes_total",
				Operator:   LessThan,
				Threshold:  1,
			},
		},
		DocLink: "./guides/002",
	},
	{
		Input: "httpjson",
		Conditions: []Condition{
			{
				MetricName: "http_response_4xx_total",
				Operator:   GreaterThanEquals,
				Threshold:  10,
			},
		},
		DocLink: "./guides/003",
	},
}
