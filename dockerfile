# Build stage
FROM golang:1.22 AS builder

# Enable Go modules and set the working directory
WORKDIR /app

COPY go.mod go.sum .env  ./
RUN go mod download

COPY . .

RUN GOOS=linux go build -o asset_price .

FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/asset_price .
COPY .env .

CMD ["./asset_price"]
