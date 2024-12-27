FROM golang:1.22-alpine as builder

RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/cmd/api

RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o /app/cmd/api/main .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/cmd/api/main .

ENV DB_HOST=""
ENV DB_PORT=""
ENV DB_USER=""
ENV DB_PASSWORD=""
ENV DB_NAME=""
ENV APP_PRODUCTION=""

EXPOSE 8080
CMD ["./main"]