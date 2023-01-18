#!/bin/bash
set -euo pipefail

dmesg

go install gotest.tools/gotestsum@latest

gotestsum --junitfile junit.xml ./...
