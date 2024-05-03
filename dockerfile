# start with base go runtime
FROM golang:1.22.2

# set working directory inside the containter
WORKDIR /k-link

# copy the go source code into the container
COPY . .

# Build the Go binary
RUN go get
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 5000

# Command to run the executable
CMD ["./main"]
