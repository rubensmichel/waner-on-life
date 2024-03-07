# use official Golang image
FROM golang:1.22.0

# set working directory
WORKDIR /app

# Copy the source code
COPY . . 

# Download and install the dependencies
RUN go get -d -v ./...


# Build the Go app
RUN go build -o app cmd/main.go

#EXPOSE the port
EXPOSE 3000

# Run the executable
CMD ["./app"]