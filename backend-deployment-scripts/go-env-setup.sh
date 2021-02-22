#!bin/bash
wget https://golang.org/dl/go1.16.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.16.linux-amd64.tar.gz
cat >> ~/.profile <<x
export PATH=$PATH:/usr/local/go/bin
x
source ~/.profile
