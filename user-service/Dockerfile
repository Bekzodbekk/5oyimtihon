FROM golang:1.22.5-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY ./config/config.yaml ./config/config.yaml

COPY --from=builder /app/main .
EXPOSE 8084
CMD ["./main"]