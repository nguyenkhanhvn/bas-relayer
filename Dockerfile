FROM golang:1.19.4-alpine as builder

RUN apk add --no-cache gcc musl-dev git

WORKDIR /go/delivery
ADD ./src .

RUN go mod tidy
RUN go build -o /build/bin/relayer ./cmd/relayer/main.go

FROM alpine:latest

COPY --from=builder /build/bin/relayer /usr/local/bin/