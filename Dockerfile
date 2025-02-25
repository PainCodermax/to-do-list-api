FROM golang:1.22-alpine3.19 AS builder

WORKDIR /app
COPY . .
RUN go build -o to-do-list-api main.go

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/to-do-list-api .
COPY .env .
RUN apk add --no-cache tzdata
EXPOSE 3000

ENTRYPOINT ["/app/to-do-list-api"]
