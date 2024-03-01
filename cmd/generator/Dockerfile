FROM golang:1.21.3-bullseye as builder

WORKDIR /opt/generator
COPY . .

RUN go build -o main ./cmd/generator

FROM debian:11.8

WORKDIR opt/generator
COPY --from=builder /opt/generator/main .

CMD ["./main"]