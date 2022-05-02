FROM golang:1.18.1 as builder
# Define build env
ENV GOOS linux
ENV CGO_ENABLED 0
# Add a work directory
WORKDIR /app
# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download
# Copy app files
COPY . .
# Build app
RUN go build -o server

FROM alpine:3.14 as production
## Add Timezone
RUN apk update && apk add tzdata
ENV TZ=Asia/Jakarta
# Add certificates
RUN apk add --no-cache ca-certificates
# Copy built binary from builder
COPY --from=builder app .
# Expose port
EXPOSE 8000
# Exec built binary
CMD ./server