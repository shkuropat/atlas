#!/bin/bash

./dev/run_code_generator.sh
./dev/format_unformatted_sources.sh
go build ./pkg/...

