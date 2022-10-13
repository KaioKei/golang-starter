FROM golang:1.19-alpine AS build
WORKDIR /
RUN mkdir -p /var/www/html/foo
COPY ../../go.mod .
COPY ../../go.mod .
RUN CGO_ENABLED=0 go get -ldflags "-s -w -extldflags '-static'" github.com/go-delve/delve/cmd/dlv
RUN CGO_ENABLED=0 go build -gcflags "all=-N -l" -o ./hello main.go

FROM scratch
COPY --from=build /go/bin/dlv /dlv
COPY --from=build /hello /hello
ENTRYPOINT [ "/dlv", "--listen=:2345", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "/hello" ]