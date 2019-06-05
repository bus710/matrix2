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
