# matrix2

## Prerequisites

- https://www.marksei.com/docker-on-raspberry-pi-raspbian/

```
sudo apt-get install \
     apt-transport-https \
     ca-certificates \
     curl \
     gnupg2 \
     software-properties-common \
     git

curl -fsSL https://download.docker.com/linux/debian/gpg | sudo apt-key add -

echo "deb [arch=armhf] https://download.docker.com/linux/debian \
     $(lsb_release -cs) stable" | \
     sudo tee /etc/apt/sources.list.d/docker.list

sudo apt-get update
sudo apt-get install docker-ce
systemctl enable --now docker
```

## How to run

```
git clone https://github.com/bus710/matrix2
cd matrix2/src/dockerARM
docker build -t matrix2 .
docker run -p 3000:3000 --device /dev/i2c-0 --device /dev/i2c-1 -it -rm --name matrix2i matrix
```

