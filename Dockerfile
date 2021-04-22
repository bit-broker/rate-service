FROM golang:1.15

# Copy source code
WORKDIR /home/rate-service
COPY . .

# Build and run
RUN go build -o rate-service main.go

# Run
CMD /home/rate-service/rate-service
