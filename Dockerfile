FROM golang:1.17-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -v -x -o /go-redis-app

FROM alpine:latest
COPY --from=builder /go-redis-app /go-redis-app
EXPOSE 8080
ENTRYPOINT ["/go-redis-app"]