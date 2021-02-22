#!/bin/bash
apt update -y

echo "\n\n########### Installing NGINX #################\n\n"

apt-get install make -y
apt-get install supervisor -y