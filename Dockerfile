FROM golang:1.23.7-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY main.go ./

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o proxy-server .

FROM alpine:3.18

WORKDIR /app

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/proxy-server .

# COPY .env ./

RUN mkdir -p /app/data

EXPOSE 8080

VOLUME ["/app/data"]

CMD ["./proxy-server"] 