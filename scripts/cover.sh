#!/bin/bash
go test $(go list ./... | grep -v  'src/templates') -coverprofile=coverage.out