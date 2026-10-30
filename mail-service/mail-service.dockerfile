FROM golang:1.26-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o mailerApp ./cmd/api
FROM alpine:latest


RUN mkdir /app

COPY --from=builder /app/mailerApp /app
COPY templates /templates

CMD ["/app/mailerApp"]