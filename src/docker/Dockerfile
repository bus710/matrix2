# Start from go-1.12.5
FROM amd64/golang:1.12.5

# Prep
RUN apt update
RUN apt install git
RUN apt-get install xz-utils -y 

# Get the repo
RUN git clone https://github.com/bus710/matrix2 /root/matrix2

# Build the backend
WORKDIR /root/matrix2/src/back
RUN go get -d -v ./...
RUN go build -o matrix2

# Run
#CMD ["./matrix2"]
