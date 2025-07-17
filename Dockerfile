# Build stage
FROM golang:1.24.5-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /main ./cmd/api

# Final stage
FROM scratch
COPY --from=builder /main /main
COPY media /media
EXPOSE 8888
ENV API_ONLY=true
CMD ["/main"]
