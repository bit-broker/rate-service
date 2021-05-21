FROM golang:1.16.4-alpine3.13

# Copy source code
WORKDIR /home/rate-service
COPY . .

# Build and run
RUN go build -o rate-service main.go

# Run
CMD /home/rate-service/rate-service
