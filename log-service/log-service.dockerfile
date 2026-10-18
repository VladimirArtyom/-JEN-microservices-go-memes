FROM golang:1.24-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o brokerApp ./cmd/api
FROM alpine:latest


run mkdir /app

COPY --from=builder /app/logApp /app

CMD ["/app/logApp"]