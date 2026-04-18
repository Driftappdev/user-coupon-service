FROM golang:1.22-alpine AS builder
WORKDIR /src
RUN apk add --no-cache ca-certificates git tzdata
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags="-s -w" -o /out/user-coupon-service ./cmd/main.go

FROM alpine:3.20
RUN apk add --no-cache ca-certificates tzdata wget && addgroup -S app && adduser -S app -G app
WORKDIR /app
COPY --from=builder /out/user-coupon-service /app/user-coupon-service
COPY config/config.yaml /app/config/config.yaml
ENV CONFIG_PATH=/app/config/config.yaml
USER app
EXPOSE 8080 50051
ENTRYPOINT ["/app/user-coupon-service"]
