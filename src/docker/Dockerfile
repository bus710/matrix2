# Start from go-1.12.5
FROM arm32v7/golang:1.12.6-alpine3.9

# Prep
RUN apk update
RUN apk add git

# Get the repo
RUN git clone https://github.com/bus710/matrix2 /root/matrix2

# Build the backend
WORKDIR /root/matrix2/src/back
RUN go get -d ./...
RUN go build -o matrix2

# Run
CMD ["./matrix2"]
