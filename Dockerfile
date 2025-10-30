# ---------- builder ----------
FROM golang:1.25-alpine AS builder
RUN apk add --no-cache git build-base

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags="-s -w" -o /app/subscriptions ./cmd/app

# ---------- runtime ----------
FROM alpine:3.18
RUN apk add --no-cache ca-certificates

RUN addgroup -S app && adduser -S -G app app

WORKDIR /app
COPY --from=builder /app/subscriptions /app/subscriptions

USER app

EXPOSE 8093
ENV PORT=8093

CMD ["/app/subscriptions"]
