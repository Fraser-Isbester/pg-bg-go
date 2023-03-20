# Build stage
FROM golang:1.17 AS builder
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o blue-green-postgres ./cmd

# Run stage
FROM gcr.io/distroless/base-debian10
COPY --from=builder /app/blue-green-postgres /
CMD ["/blue-green-postgres"]