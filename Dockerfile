# Use lightweight Alpine base image
FROM golang:1.24-alpine


# Set the working directory inside the container
WORKDIR /app

# Copy all files from the project directory to the container
COPY . .

# Install necessary dependencies (if any)
RUN apk add --no-cache make

RUN go mod tidy

RUN make install_swagger

RUN make docs

RUN go build -o microecom .

# Expose the port your app runs on
EXPOSE 9090

# Command to run the Go app
CMD ["./microecom"]
