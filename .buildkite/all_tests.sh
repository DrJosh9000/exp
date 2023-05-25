#!/bin/bash
set -euo pipefail

go install gotest.tools/gotestsum@latest
gotestsum --junitfile junit.xml ./...

