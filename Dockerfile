
FROM golang:1.23.2 AS builder
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /build
COPY . .
RUN go build -o main cmd/ordersystem/main.go cmd/ordersystem/wire_gen.go \
  && chmod +x main

FROM scratch
WORKDIR /app
COPY --from=builder /build/main /app/main
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY .env ./
EXPOSE 8000 50051 8080
CMD ["/app/main"]
