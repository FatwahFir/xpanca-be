# ---- Build stage
FROM golang:1.23-alpine AS builder
WORKDIR /app
RUN apk add --no-cache git ca-certificates
COPY go.mod go.sum* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/api ./cmd/api
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/seed ./cmd/seed

# ---- Runtime stage
FROM gcr.io/distroless/base-debian12:nonroot
WORKDIR /
COPY --from=builder /bin/api /bin/api
COPY --from=builder /bin/seed /bin/seed
COPY .env .env
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/bin/api"]
