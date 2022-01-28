FROM golang:1.17-alpine AS builder

RUN mkdir /app
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY main.go .
RUN go build main.go

FROM registry.nordix.org/cloud-native/kpng:latest

COPY --from=builder /app/main /main
COPY entrypoint.sh /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
