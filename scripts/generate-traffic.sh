#!/usr/bin/env bash

URL="${1:-https://demo-api.demo.local}"

while true; do
    curl -k -s "$URL"
    sleep 0.1
done