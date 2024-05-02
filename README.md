# GOLANG TUTORIALS

## Sources

- [Project layout standards](https://github.com/golang-standards/project-layout)
- [Best Practices](https://github.com/golang/go/wiki/CodeReviewComments)
- [Official tutorials](https://pkg.go.dev/cmd/go)
- [Intermediate tutorials](https://github.com/callicoder/golang-tutorials)

## Libs

* Cli :
    * Cobra
* Logging :
    * zap
* Files Mapping :
    * Viper
* ORM :
    * Gorm
* API :
    * REST :
        * Gin

## Install commands

```sh
./install.sh
```

## Init a module

```sh
# will create a go.mod file in ./
go mod init my_module_name
```

## Add a new module

```sh
# import
go get <module import>
# synchronize code's dependencies
go mod tidy
# if vendors
go mod vendor
```

## Build Manually

```sh
# example for hello command
cd cmd/hello
go build
./hello
```

## Testing one file

Files named `impl1.go`, `utils.go` etc .. are implementation files containing required function and vars definition for
the test

```sh
go test impl1_test.go impl1.go utils.go
```

## Remote debug

You can use the `hello-debug` command to debug remotely.

Install Delve :

```sh
go install github.com/go-delve/delve/cmd/dlv@latest
# before 1.18 :
go get github.com/go-delve/delve/cmd/dlv@latest
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

## Protocol buffer

Doc from [gRPC Quickstart](https://grpc.io/docs/languages/go/quickstart/).

Install on your system :

```sh
sudo apt install -y protobuf-compiler
protoc --version
```

Install the plugins for Golang :

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

Compile a proto file :

```sh
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld/helloworld.proto
```

## Private dependency with authentication

In case you need to get a module from a private repo with https auth, you need to configure the authentication in
the `.git-credentials` file and configure git to use this file.

## Get development branch of a module

```sh
# if branch name is 'aes-cbc-support'
go get github.com/IceManGreen/crypto11@aes-cbc-support                                                                            [29/534]
> go: downloading github.com/pkg/errors v0.9.1
> go: added github.com/IceManGreen/crypto11 v0.0.0-20240118165436-b187874cdf68
go mod tidy
```

## Replace a dependency by another

1. Remove vendor folder and mod cache :

```sh
cd mygoproject/
rm -rf vendor/ go.sum
go clean -modcache
```

2. Remove all reference of the old dependency in `go.mod`.
3. Replace in all `*.go` files of the project the old dependency by the new one
4. Get the new dependency with `GOPROXY=direct go get -u -v github.com/group/repo`. For a dev dependency, you can directly
   type `GOPROXY=direct go get -uv github.com/group/repo@branch`
5. Run `go mod tidy` and `go mod vendor`

## Force update dependency on newer commits

```sh
# on default branch :
GOPROXY=direct go get -u github.com/IceManGreen/gose
# on a specific branch :
GOPROXY=direct go get -u github.com/IceManGreen/gose@aes-cbc-support 
```