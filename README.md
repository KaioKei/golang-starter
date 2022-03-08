# GOLANG STARTER

## Download

Install this project into `$GO_ROOT/src/`.

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

Manually :

```sh
cd cmd/hello
go build
./hello
```

From custom script:

```sh
./install.sh
./bin/hello
```
