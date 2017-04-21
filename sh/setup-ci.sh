#!/usr/bin/env bash


# install plz
curl https://get.please.build | bash

# install go
sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt-get update
sudo apt-get install golang-go -y
sudo apt-get install build-essential unzip -y
pushd /tmp
wget https://github.com/google/protobuf/releases/download/v3.2.0/protoc-3.2.0-linux-x86_64.zip
unzip protoc-3.2.0-linux-x86_64.zip
mv /tmp/bin/protoc /usr/local/bin/
popd