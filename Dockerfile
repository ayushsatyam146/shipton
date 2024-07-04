# Use the official Golang image as a base image to build the binary
FROM golang:1.18 as builder

# Set the Current Working Directory inside the container
WORKDIR /workspace

# Copy go.mod and go.sum files to the workspace
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the workspace
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o shipton-controller main.go

# Use a minimal image as a base image
FROM gcr.io/distroless/static:nonroot

# Copy the Shipton controller binary from the builder stage
COPY --from=builder /workspace/shipton-controller /shipton-controller

# Command to run the binary
ENTRYPOINT ["/shipton-controller"]
