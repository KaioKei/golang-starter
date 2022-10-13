#cp Dockerfile ../../

FROM golang:1.19-alpine AS build

ENV PKG golang_starter/pkg/greetings

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
COPY cmd/hello-debug ./
COPY pkg/greetings /usr/local/go/src/$PKG

RUN go mod download

RUN CGO_ENABLED=0 go install github.com/go-delve/delve/cmd/dlv@latest
RUN CGO_ENABLED=0 go build -gcflags="all=-N -l" -o /app/hello ./main.go

FROM scratch
COPY --from=build /go/bin/dlv /dlv
COPY --from=build /app/hello /hello
#/dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient exec /hello
EXPOSE 2345
ENTRYPOINT [ "/dlv", "--listen=:2345", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "/hello" ]