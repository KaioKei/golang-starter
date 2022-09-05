# GOLANG STARTER

## Download

Install this project into `$GO_ROOT/src/` or create a symbolic link to it.

```sh
git clone https://gitlab.com/MrCachou/golang_starter
```

## Development

Use goenv locally with this project.

## Init project

```sh
mkdir myproject && cd myproject
goenv local 1.19.0
go mod init myproject
tree -a
.
├── go.mod
└── .go-version
```

## Project architecture

```sh
tree -a -L 1 
.
├── cmd
├── go.mod
├── .go-version
├── internal
├── pkg
├── README.md
└── test
```

## Test

```sh
# x/... matches x as well as x's subdirectories.
go test ./test/...
```

## Run

```sh
go run cmd/hello/main.go
```

## Install

```sh
./install.sh
./bin/hello
```

Or manually :

```sh
# install ..
export GOBIN="$(realpath ./bin)"
go install ./cmd/hello
# .. or build
go build ./cmd/hello -o ./cmd/hello/hello
# then run
./bin/hello
```

## Container Remote debug

Using Delve, Buildah and Podman.

Install Delve :

```sh
go install github.com/go-delve/delve/cmd/dlv@latest
```

Build or install your app with debug flags :

```sh
# build ... 
go build -gcflags="all=-N -l" -o /tmp/hello-debug ./cmd/hello-debug 
```

```sh
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
buildah images  # should show localhost/hello-debug
```

Run the container with podman and expose the Delve's debug port : 

```sh
podman run -p 2345:2345 --name hello-debug --security-opt seccomp=unconfined localhost/hello-debug
```

Then you can use your IDE for remote debug using the same host:port to communicate with the application.

**THE OPTION `--security-opt seccomp=unconfined` IS MANDATORY TO ALLOW DELVE TO RUN INSIDE THE CONTAINER**

### Kubernetes Remote Debug

Firts, build the image following the [Container Remote Debug section](#container-remote-debug).

Then push the image in a local registry :

```sh
podman tag localhost/hello-debug localhost:5000/hello-debug
podman push localhost:5000/hello-debug
```

    **WARN**: How to create a private registry is out of scope 

Create a debug pod :

```yaml
---
apiVersion: v1
kind: Pod
metadata:
  name: hello-debug
  namespace: test
spec:
  containers:
    - name: hello-debug
      image: localhost:5000/hello-debug
      imagePullPolicy: Always
      ports:
        - containerPort: 2345
```

Forward the debug port :

```sh
kubectl -n test port-forward pod/hello-debug 2345:2345
```

And start debugging !

You also can print the logs to help during debug :

```sh
kubectl -n test logs hello-debug -f
```

## Sources

1.[Basic Golang official tutorial: hello](https://go.dev/doc/tutorial/getting-started)
2.[Golang thematical tutorial from Callicoder](https://github.com/callicoder/golang-tutorials)
