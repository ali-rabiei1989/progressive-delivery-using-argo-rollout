# Traffic Simulation

## Purpose

The request generator creates enough HTTP requests for Prometheus to collect meaningful metrics during canary analysis.

The script is intentionally simple and has no external dependency beyond `curl`.

## Script

```bash
#!/usr/bin/env bash

URL="${1:-https://demo-api.demo.local}"

while true; do
    for i in {1..100}; do
        curl -k -s -o /dev/null "$URL" &
    done

    wait
    sleep 1
done
```

## Usage

```bash
chmod +x scripts/request-generator.sh
./scripts/request-generator.sh https://demo-api.demo.local
```

## Why It Is Needed

Prometheus calculates rates over time windows. Without traffic, the application metrics may not contain enough data for the AnalysisTemplate to evaluate the rollout.
