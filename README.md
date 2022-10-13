# GOLANG TUTORIALS

## Sources

Common standard : https://github.com/golang-standards/project-layout
Best Practices : https://github.com/golang/go/wiki/CodeReviewComments
Official tutorials : https://pkg.go.dev/cmd/go
Intermediate tutorials : https://github.com/callicoder/golang-tutorials

## Libs

* Cli :
  * Cobra
* Logging :
  * zap
* Files Mapping :
  * Viper
* ORM :
  * Gorm

## Add a new module

```sh
go get <module import>
go mod tidy
```

## Remote debug

You can use the `hello-debug` command to debug remotely.

Install Delve :

```sh
go install github.com/go-delve/delve/cmd/dlv@latest
```

Build or install your app with debug flags :

```sh
# build ...                                                                                 
go build -gcflags="all=-N -l" cmd/hello-debug -o /tmp/hello-debug
# ... or install
export GOBIN=/tmp
go install -gcflags="all=-N -l" ./cmd/hello-debug  # output in /tmp
```

Run the application with Delve to debug :

```sh
dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient exec /tmp/hello-debug     
```

Then you can use your IDE for remote debug using the same host:port to communicate with the application.  
**Do not forget to expose the port if you want to debug inside a container !**

## Remote debug in container

You can use the `hello-debug` command to debug remotely :

1. `cp cmd/hello-debug/Dockerfile ./`
2. `podman build -t hello-debug .`
3. `podman run -d --name hello-debug -p 2345:2345 --security-opt label=disable --privileged localhost/hello-debug`. From
here, the code in the container won't start before the connexion is established with the debug client. You can check
this behavior using :

```sh
$ podman logs -f hello-debug
API server listening at: [::]:2345
2022-10-13T15:05:35Z warning layer=rpc Listening for remote connections (connections are not authenticated nor encrypted)
```

4. Put a breakpoint in the code and create a Go Remote Debug configuration in the Intellij Runner to `localhost:2345`.
