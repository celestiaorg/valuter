#!/bin/bash

set -o errexit -o nounset

# Read common EVN vars
export $(cat ../.env | sed '/^#/d' | xargs)

export VALUTER_SERVING_ADDR=:8080

go mod tidy  && reset && go build -o app . && ./app
