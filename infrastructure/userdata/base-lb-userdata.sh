#!/bin/bash
apt-get update -y

echo "\n\n########### Installing NGINX #################\n\n"
apt-get install nginx -y
systemctl start nginx