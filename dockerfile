#------------------------------------------------------------------
FROM golang:1.17-alpine as builder

# Create and change to the 'code' directory.
WORKDIR /code

# Install dependencies.
COPY go.mod .
COPY go.sum .
RUN go mod download

# Build application binary.
COPY . .
RUN GOOS=linux go build -tags musl -v -o bin/main

#-------------------------------------------------------------------
FROM alpine:3

# Required if application uses TLS.
RUN apk add ca-certificates

# Create and change to the the bin directory.
WORKDIR /usr/bin

# Copy the files to the production image from the builder stage.
COPY --from=builder /code/bin .

# Run the web service on container startup.
CMD ["main"]

#-------------------------------------------------------------------
