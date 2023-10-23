package main

import (
	"archive/zip"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"regexp"
	"strings"
)

var (
	diag string
)

func init() {
	flag.StringVar(&diag, "diag", "", "agent diagnostics zip file")
}

func main() {
	log.SetFlags(0)
	flag.Parse()

	if diag == "" {
		log.Fatal("-diag is required")
	}

	zipListing, err := zip.OpenReader(diag)
	if err != nil {
		log.Fatal(err)
	}
	defer zipListing.Close()

	filepath_metrics_re, err := regexp.Compile("components/.+/input_metrics.json")
	if err != nil {
		return
	}

	var inputMetrics []map[string]interface{}

	for _, file := range zipListing.File {
		match := filepath_metrics_re.MatchString(file.Name)

		if match {
			fmt.Printf("Analyzing %s\n", file.Name)

			reader, err := file.Open()
			if err != nil {
				log.Fatal("Failed to read input file.", err)
			}
			defer reader.Close()

			dec := json.NewDecoder(reader)
			for {
				var result []map[string]interface{}
				if err := dec.Decode(&result); err == io.EOF {
					break
				} else if err != nil {
					log.Fatal(err)
				}

				if len(result) == 0 {
					// sometimes we have empty metrics files
					continue
				}

				inputMetrics = append(inputMetrics, result...)
			}
		}
	}

	var metricWarnings []MetricWarning

	for _, metricset := range inputMetrics {
		for _, check := range checks {
			if metricset["input"] == check.Input {
				matchCount := 0
				for _, condition := range check.Conditions {

					value := getMetricValue(condition.MetricName, metricset)
					var threshold float64 = float64(condition.Threshold)

					if condition.Operator == GreaterThan && value > threshold ||
						condition.Operator == GreaterThanEquals && value >= threshold ||
						condition.Operator == LessThan && value < threshold ||
						condition.Operator == LessThanEquals && value <= threshold {
						matchCount++
					}
				}

				if matchCount == len(check.Conditions) {
					warning := MetricWarning{
						InputName: metricset["id"].(string),
						Check:     check,
					}
					metricWarnings = append(metricWarnings, warning)
				}
			}
		}
	}

	if len(metricWarnings) > 0 {
		for _, warning := range metricWarnings {
			fmt.Println(warning)
		}
	} else {
		fmt.Println("no warnings identified")
	}
}

func getMetricValue(metricName string, metricset map[string]interface{}) float64 {
	var value float64
	metricNameParts := strings.Split(metricName, ".")

	if len(metricNameParts) > 1 {
		// histogram metric
		metric := metricset[metricNameParts[0]]

		jsonBytes, _ := json.Marshal(metric)
		dec := json.NewDecoder(strings.NewReader(string(jsonBytes)))
		for {
			var hm HistogramMetric
			if err := dec.Decode(&hm); err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}
			value = hm[metricNameParts[1]][metricNameParts[2]]
		}
	} else {
		// numerical metric
		if nm, ok := metricset[metricName].(float64); ok {
			value = nm
		} else {
			log.Fatal("unable to parse numerical metric")
		}
	}
	return value
}
