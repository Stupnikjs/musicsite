# Stage 2 - Your application image
FROM golang:1.22-alpine
WORKDIR /app


COPY go.mod .
COPY go.sum .

COPY . .

# google credentials.json

# Download all dependencies
RUN go mod download
# Copy the source code into the container


 

# Build the Go application
RUN go build -o main ./cmd
# Specify the command to run when the container starts

EXPOSE 8080


CMD [ "./main" ]