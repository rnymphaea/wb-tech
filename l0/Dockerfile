FROM golang:1.24.4-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o server ./cmd/main.go

FROM alpine:3.19

RUN adduser -D appuser
WORKDIR /home/appuser
COPY --from=builder /app/server .
USER appuser

EXPOSE 8080
CMD ["./server"]

