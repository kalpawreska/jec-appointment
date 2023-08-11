# Base Image aka Builder
# FROM golang:alpine
FROM golang AS builder

# Working directory
WORKDIR /app

# Copy all files
COPY . .

# Running command
# RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/grpc/main.go
# RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main ./cmd/api/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main ./cmd/grpc/main.go

# Base Image aka alpine
# FROM alpine

# Base Image aka scratch
FROM scratch

# Install CA-Certificate
RUN  opkg-install ca-certificates

# Working directory
WORKDIR /app

# Copy all files from builder
COPY --from=builder /app/main .

# Application Environtment
ENV APP_PORT=":9099"

# Install CA-Certificate
EXPOSE 443

# Run command
CMD [ "./main"]

# Base Image aka Builder
# FROM golang:alpine
FROM golang AS builder

# Working directory
WORKDIR /app

# Copy all files
COPY . .

# Running command
# RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/grpc/main.go
# RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main ./cmd/api/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main ./cmd/grpc/main.go

# Base Image aka alpine
# FROM alpine

# Base Image aka scratch
FROM scratch

# Install CA-Certificate
RUN  opkg-install ca-certificates

# Working directory
WORKDIR /app

# Copy all files from builder
COPY --from=builder /app/main .

# Application Environtment
ENV APP_PORT=":9099"

# Install CA-Certificate
EXPOSE 443

# Run command
CMD [ "./main"]
