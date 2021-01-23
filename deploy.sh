#!/bin/sh
if ! type docker >/dev/null 2>&1
then 
    sudo apt-get remove docker docker-engine docker.io
    sudo apt-get install apt-transport-https ca-certificates curl gnupg-agent software-properties-common
    sudo apt-get update
    $ curl -fsSL https://mirrors.aliyun.com/docker-ce/linux/ubuntu/gpg | sudo apt-key add -
    sudo add-apt-repository \
    "deb [arch=amd64] https://mirrors.aliyun.com/docker-ce/linux/ubuntu \
    $(lsb_release -cs) \
    stable"
    sudo apt-get update
    sudo apt-get install docker-ce docker-ce-cli containerd.io
    sudo systemctl enable docker
    sudo systemctl start docker
else
    sudo apt-get update
    echo 'test jenkins'
    echo 'hello world'
fi