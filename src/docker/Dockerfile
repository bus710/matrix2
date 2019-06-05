# Start from go-1.12.5
FROM amd64/golang:1.12.5

# Prep
RUN apt update
RUN apt install git
RUN apt-get install xz-utils -y 

# Get the repo
RUN git clone https://github.com/bus710/matrix2 /root/matrix2

# Get flutter SDK
WORKDIR /root
RUN wget https://storage.googleapis.com/flutter_infra/releases/stable/linux/flutter_linux_v1.5.4-hotfix.2-stable.tar.xz
RUN tar xf flutter_linux_v1.5.4-hotfix.2-stable.tar.xz
ENV PATH="${PATH}:/root/flutter/bin"
ENV PATH="${PATH}:/root/flutter/bin/cache/dart-sdk/bin"
ENV PATH="${PATH}:/root/flutter/.pub-cache/bin"
RUN flutter channel master
RUN flutter precache
RUN flutter pub global activate webdev
RUN flutter pub global activate stagehand

# Build the frontend
WORKDIR /root/matrix2/src/front
RUN flutter pub get
RUN webdev build --release

# Build the backend
WORKDIR /root/matrix2/src/back
RUN go get -d -v ./...
RUN go build -o matrix2

# Run
CMD ["./matrix2"]
