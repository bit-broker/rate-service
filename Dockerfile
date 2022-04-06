FROM golang:1.17.3-alpine3.13

# Add build base
RUN apk add build-base

# Copy source code
WORKDIR /home/rate-service
COPY . .

# Build and run
RUN go build -o rate-service main.go

# Run
CMD /home/rate-service/rate-service
