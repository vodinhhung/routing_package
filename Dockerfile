FROM ubuntu:latest
LABEL authors="vohung"

ENTRYPOINT ["top", "-b"]

FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
RUN go build -o server

FROM alpine
WORKDIR /app
COPY --from=builder /app/server .
CMD ["./server"]
