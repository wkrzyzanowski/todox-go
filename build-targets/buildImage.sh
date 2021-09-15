#!/bin/bash
set -o errexit  # Force to exist on error
set -o errtrace # Print error trace
set -o nounset  # Exit when variables are not initialized
#set -o xtrace # Print debug for all executed commands

gitroot=$(git rev-parse --show-toplevel)
source "${gitroot}/build-targets/common/printTools.sh"

printInfo "Build Docker image for ToDoX application"

set -x
docker build --rm -t "todox-go:1.0.0" -f "${gitroot}/build-targets/container/Dockerfile" "${gitroot}"
set +x

printInfo "Docker image was built successfully."
