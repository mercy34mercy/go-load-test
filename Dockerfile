# Build stage
FROM golang:1.22-alpine as builder

WORKDIR /app/src

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o /app/main .

# Runtime stage
FROM gcr.io/distroless/base-debian10

COPY --from=builder /app/main .

CMD ["/main"]