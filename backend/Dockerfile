# Use the official Go image as the build stage
FROM golang:1.25.4-alpine3.22 AS builder

WORKDIR /app

COPY backend/go.mod backend/go.sum ./

RUN go mod download

COPY backend/ .

RUN go build -o tempo_time_tracker ./cmd/backend

FROM alpine:3.22.2

WORKDIR /opt/tempo_time_tracker

COPY --from=builder /app/tempo_time_tracker .

EXPOSE 8080

CMD ["./tempo_time_tracker"]
