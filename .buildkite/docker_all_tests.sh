#!/bin/bash
set -euo pipefail

docker run -t -i --rm --init \
  --volume ${PWD}:/workdir \
  --workdir /workdir \
  --runtime runsc \
  --label com.buildkite.job-id=${BUILDKITE_JOB_ID} \
  golang:latest /bin/sh -e -c ./.buildkite/all_tests.sh