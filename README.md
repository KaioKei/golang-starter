# GOLANG STARTER

## Download

Install this project into `$GO_ROOT/src/` or create a symbolic link to it.

## Development

Use goenv locally with this project.

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
```

Run the container with podman and expose the Delve's debug port : 

```sh
podman run -p 2345:2345 --name hello-debug --security-opt seccomp=unconfined localhost/hello-debug
```

Then you can use your IDE for remote debug using the same host:port to communicate with the application.

**THE OPTION `--security-opt seccomp=unconfined` IS MANDATORY TO ALLOW DELVE TO RUN INSIDE THE CONTAINER**

### Kubernetes Remote Debug

Firts, build the image following the [Container Remote Debug section](#container-remote-debug).

Then push the image in a local registry.

Create a deployment :

```yaml
---
apiVersion: v1
kind: Namespace
metadata:
  name: test
---        
apiVersion: apps/v1     
kind: Deployment        
metadata:  
  labels:  
    app: hello-debug-app
  name: hello-debug     
  namespace: test      
spec:      
  replicas: 3           
  selector:
    matchLabels:        
      app: hello-debug-app           
  template:
    metadata:           
      labels:           
        app: hello-debug-app         
    spec:  
      containers:       
        - image: localhost:5000/hello-debug   
          name: hello-debug
          ports:        
            - containerPort: 2345      
---        
apiVersion: v1          
kind: Service           
metadata:  
  name: hello-debug-svc 
  namespace: hello      
  labels:  
    app: hello-debug-app
spec:      
  selector:
    app: hello-debug-app
  ports:   
    - protocol: TCP     
      port: 2345
      targetPort: 2345
```

## Sources

1.[Basic Golang official tutorial: hello](https://go.dev/doc/tutorial/getting-started)
2.[Golang thematical tutorial from Callicoder](https://github.com/callicoder/golang-tutorials)
