FROM golang:1.21-alpine

# Add environment variables
ENV KAFKA_HOST=12.12.12.12 \
    KAFKA_PORT=9092

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY ${PWD} /app    
RUN go mod download

# Build the binary.
RUN go build -o /denv-register-producer

# Run backend on port 8080
EXPOSE 8080

# Run the web service on container startup
CMD ["/denv-register-producer"]