#! /bin/bash

sudo apt-get update
sudo apt-get install \
    apt-transport-https \
    ca-certificates \
    curl \
    lsb-release \
    software-properties-common \
    nano

# Get Docker
echo "Checking for Docker"
if ! [ -x "$(command -v docker)" ]; then
    echo "Fetching docker"
    curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg

    echo \
    "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu \
    $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

    sudo apt-get update
    sudo apt-get install docker-ce docker-ce-cli containerd.io
fi
echo "We have docker"
echo docker -v

# Get docker-compose
echo "Checking for docker-compose"
if ! [ -x "$(command -v docker-compose)" ]; then
    echo "Fetching docker-compose"
    sudo curl -L "https://github.com/docker/compose/releases/download/1.23.1/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
    sudo chmod +x /usr/local/bin/docker-compose
fi
echo "We have docker-compose"
echo docker-compose -v

docker-compose up -d
