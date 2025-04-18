#!/bin/bash
go mod edit -go=$1 -print
go mod edit -go=$1
go mod tidy
go build -o "${2:-linkslasher}"