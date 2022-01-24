FROM golang:1.16.5 as development
# Add a work directory
WORKDIR /app
# Cache and install dependencies
COPY go.mod ./
RUN go mod download
# Copy app files
COPY . .
# Install Reflex for development
RUN go install github.com/cespare/reflex@latest
# Expose port
EXPOSE 8080
# Start app
CMD reflex -g '*.go' go run main.go --start-service