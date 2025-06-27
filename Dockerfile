# ---- Build Stage ----
# Use an official Go image that matches your local environment.
FROM golang:1.24-alpine AS builder

# Set the working directory inside the container.
WORKDIR /app

# Copy all project files (go.mod, go.sum, main.go) into the container.
# Using 'COPY . .' is a robust way to bring in all source files.
COPY . .

# The 'go mod tidy' command ensures all necessary dependencies are present
# and removes any unused ones.
RUN go mod tidy

# Build the Go application.
# CGO_ENABLED=0 is important for creating a static binary.
# -ldflags "-w -s" strips debugging info, reducing the binary size.
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/word-splitter -ldflags "-w -s" .

# ---- Final Stage ----
# Use the minimal 'scratch' base image for the final container.
# It contains nothing but our application binary.
FROM scratch

# Set the working directory.
WORKDIR /app

# Copy the built binary from the builder stage.
COPY --from=builder /app/word-splitter /app/word-splitter

# Set the entrypoint for the container. This is the command that runs.
ENTRYPOINT ["/app/word-splitter"]
