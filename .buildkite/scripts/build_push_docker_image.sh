#!/bin/bash

set -euo pipefail

source .buildkite/scripts/common.sh

echo "Building the docker image..."
if ! docker pull -q ${DOCKER_IMAGE}:${DOCKER_IMAGE_SHA_TAG} 2> /dev/null; then
    DOCKER_IMAGE_TAG="${DOCKER_IMAGE_SHA_TAG}"
    DOCKER_IMAGE=${DOCKER_IMAGE} DOCKER_IMAGE_TAG=${DOCKER_IMAGE_TAG} make build-and-push-docker
fi

if [ -n "${BUILDKITE_TAG}" ]; then
    DOCKER_IMAGE=${DOCKER_IMAGE} DOCKER_IMAGE_TAG="${DOCKER_IMAGE_GIT_TAG}" make build-and-push-docker
else
    DOCKER_IMAGE=${DOCKER_IMAGE} DOCKER_IMAGE_TAG="${DOCKER_IMAGE_LATEST_TAG}" make build-and-push-docker
fi
