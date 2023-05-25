#!/bin/bash
set -euo pipefail

buildkite-agent lock acquire LLAMA
trap "buildkite-agent lock release LLAMA" EXIT

go install gotest.tools/gotestsum@latest
gotestsum --junitfile junit.xml ./...

