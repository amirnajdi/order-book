# Build Stage
FROM golang:1.21rc2-alpine AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /order-book \
    && CGO_ENABLED=0 GOOS=linux go build -o /insert-test-data -C Test

# Production Stage
FROM alpine:latest

WORKDIR /

COPY --from=build-stage /order-book /order-book
COPY --from=build-stage /insert-test-data /insert-test-data

CMD ["/order-book"]
