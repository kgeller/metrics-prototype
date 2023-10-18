# metrics-prototype

This tool aims to help identify potential issues with Agent integrations by analyzing filebeat input metrics. It will scan an elastic agent diagnostic for the included `input-metrics.json` files, and print out identified warning conditions.

## Stack requirements

Stack version must be > 8.10 or the diagnostic will not include metrics.

## How to run

`go run github.com/kgeller/metrics-prototype@main -diag=elastic-agent-diagnostics-2023-09-26T15-45-32Z-00.zip`

## Example output

```
Analyzing components/filestream-monitoring/input_metrics.json
Analyzing components/log-default/input_metrics.json
Analyzing components/udp-default/input_metrics.json
[{udp-panw.panos-091a9b72-50e5-40df-9bad-e0b8ce13b07c {udp ./guides/002 [{received_bytes_total lt 1}]}}]
```