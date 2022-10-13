#!/usr/bin/env bash

# build first with :
#go build -gcflags="all=-N -l" -o /tmp/hello-debug ./cmd/hello-debug

IMAGE="registry.access.redhat.com/ubi8/ubi"
CONTAINER=$(buildah from "$IMAGE")

# buildah copy container_name source_host destination_container
# copy the built go app to debug
buildah copy "$CONTAINER" /tmp/hello-debug /usr/bin/hello-debug
# copy delve binary to launch the debugging over the application
buildah copy "$CONTAINER" $(which dlv) /usr/bin/dlv
# configure the container to expose the port 2345
# configure the launch command to start the application with delve with the proper debuggin port
buildah config \
    --port 2345 \
    --cmd "/usr/bin/dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient exec /usr/bin/hello-debug" \
    "$CONTAINER"
# commit the container, i.e create a new image with the configured container
buildah commit "$CONTAINER" "localhost/hello-debug"