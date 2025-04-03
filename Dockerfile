# Build stage
FROM golang:1.24.2-alpine3.21 AS builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app/main .

# Run stage
FROM alpine:3.21
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .

EXPOSE 8080
CMD [ "/app/main" ]

# Use the following command to build the Docker image
# docker build -t go-web-app .
# Use the following command to run the Docker container
# docker run -p 8080:8080 go-web-app