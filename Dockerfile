FROM golang:1.21 AS builder

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o main .

FROM alpine:latest


RUN adduser -D appuser

USER appuser

WORKDIR /home/appuser

COPY --from=builder /build/main .

EXPOSE 8080

CMD ["./main"]
