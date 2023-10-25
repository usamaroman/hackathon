FROM golang:alpine as builder
ENV POSTGRES_HOST="" \
    POSTGRES_PORT="" \
    POSTGRES_DB="" \
    POSTGRES_USER="" \
    POSTGRES_PASSWORD=""

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy -e
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -mod=mod -a -o bin cmd/main/main.go

FROM scratch
WORKDIR /app
COPY --from=builder /app/bin .
EXPOSE 8000
CMD ["/app/bin"]