FROM golang:1.20-alpine3.17 AS builder

WORKDIR /app

COPY . /app/

RUN go build -o main .

FROM alpine

WORKDIR /app

COPY --from=builder /app/main /app/main

EXPOSE 3000

CMD ["/app/main"]