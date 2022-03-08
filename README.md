# GOLANG STARTER

## Download

Install this project into `$GO_ROOT/src/` or create a symbolic link to it.

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

## Sources

1.[Basic Golang official tutorial: hello](https://go.dev/doc/tutorial/getting-started)
2.[Golang thematical tutorial from Callicoder](https://github.com/callicoder/golang-tutorials)
