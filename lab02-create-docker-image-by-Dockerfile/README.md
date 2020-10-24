# Create Docker Image by Dockerfile

Dockerfile is a text file that contains instructions for building an image.

## Build image with Dockerfile Tag 0.0.3

```sh
docker build -t <name>/hello-world:0.0.3
```

```sh
docker image ls
```

```dockerfile
FROM golang:1.15.2-alpine3.12 AS builder
WORKDIR /go/src/app
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hello .

FROM alpine:3.12
WORKDIR /root/
COPY --from=builder /go/src/app/hello .
CMD ["./hello"]
```

## Run Your Own Image Tag 0.0.3

```sh
docker container run <name>/hello-world:0.0.3
```

## Modify hello-world

> Replace this line

```go
    fmt.Println("Hello, World!")
```

> with

```go
fmt.Println("Hello, World!", time.Now())
```

## Build image with Dockerfile Tag 0.0.4

```sh
docker build -t <name>/hello-world:0.0.4
```

## Run Your Own Image Tag 0.0.4

```sh
docker container run <name>/hello-world:0.0.4
```

## Push Your Image

```sh
docker image push <name>/hello-world:0.0.4
```

## Tag latest

> docker image tag <name>/hello-world:0.0.4 <name>/hello-world:latest
```sh
docker image tag boyone/hello-world:0.0.4 boyone/hello-world:latest
```
