# matrix2

## Prerequisites

Enable i2c communication (in 5 - Interfacing Options)

```
sudo raspi-config
```

Get Docker 

```
sudo apt-get install \
     apt-transport-https \
     ca-certificates \
     curl \
     gnupg2 \
     software-properties-common \
     git

curl -sSL https://get.docker.com | sh

sudo systemctl enable docker
sudo systemctl start docker
```

Add the user to the docker group in /etc/group

```
sudo vi /etc/group
```

Don't forget rebooting

```
sudo reboot
```

## How to run

```
git clone https://github.com/bus710/matrix2
cd matrix2/src/docker
docker build -t matrix2 .
docker run -p 3000:3000 --device /dev/i2c-1 -it --rm --name matrix2i matrix2
```

