FROM golang:1.22 AS builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/main .

FROM debian:bullseye

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/main /app/main
COPY .env .

RUN ls -la /app
RUN ls -la /app/main

EXPOSE 8080

CMD ["/app/main"]