#!/usr/bin/env bash

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/Employee-Canonical-Demo-Mapper .
docker build . -t docker.repo1.uhc.com/hemi-rxclaim/employee_canonical_demo_mappe:local
