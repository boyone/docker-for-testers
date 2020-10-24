FROM golang:1.15.2-alpine3.12 AS builder
WORKDIR /go/src/app
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:3.12 
WORKDIR /root/
COPY --from=builder /go/src/app/app .
CMD ["./app"] 