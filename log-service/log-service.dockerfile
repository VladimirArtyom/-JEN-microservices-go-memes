FROM alpine:latest

run mkdir /app

COPY logApp /app

CMD ["/app/logApp"]