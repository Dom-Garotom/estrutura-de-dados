FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN go build -o /blacklist .

FROM alpine:3.22

WORKDIR /app

COPY --from=builder /blacklist /app/blacklist

CMD ["/app/blacklist"]
