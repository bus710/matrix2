FROM amd64/golang:1.12.5

RUN apt update
RUN apt install git

RUN git clone https://github.com/bus710/matrix2 /root/matrix2

WORKDIR /root/matrix2/src/back
RUN go get -d -v ./...
RUN go build -o matrix2

CMD ["./matrix2"]
