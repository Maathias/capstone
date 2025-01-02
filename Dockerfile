# builder
FROM golang:1.23.3 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

# runner
FROM registry.access.redhat.com/ubi9/ubi-minimal:latest

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]