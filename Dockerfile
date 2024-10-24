# STAGE 1 -- Build the app locally
FROM golang:1.20-alpine AS builder

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o server ./cmd/server

# STAGE 2 -- Build the image
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/server .
# Create the upload dir
RUN mkdir -p /root/uploaded_files
EXPOSE 8080
CMD ["./server"]
