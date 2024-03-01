FROM golang:1.21.3-bullseye as builder

WORKDIR /opt/app
COPY . .

RUN go build -o main ./cmd/app

FROM debian:11.8

EXPOSE 8080
WORKDIR opt/app
COPY --from=builder /opt/app/main .

CMD ["./main"]