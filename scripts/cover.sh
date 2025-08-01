#!/bin/bash
cd ..
COVERAGE=$(go tool cover -func=cover.out | grep total | awk '{print substr($3, 1, length($3)-1)}')
 echo $COVERAGE
COLOR="orange"
if (( $(echo "$COVERAGE <= 50" | bc -l) )); then COLOR="red"; elif (( $(echo "$COVERAGE > 80" | bc -l) )); then COLOR="green"; fi
echo COLOR
curl "https://img.shields.io/badge/coverage-$COVERAGE%25-$COLOR" > images/badge.svg