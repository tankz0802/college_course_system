#!/bin/sh
######### 安装docker ###########
if ! type docker >/dev/null 2>&1
then
    apt-get remove docker docker-engine docker.io
    apt-get install apt-transport-https ca-certificates curl gnupg-agent software-properties-common
    apt-get update
    $ curl -fsSL https://mirrors.aliyun.com/docker-ce/linux/ubuntu/gpg |  apt-key add -
    add-apt-repository \
    "deb [arch=amd64] https://mirrors.aliyun.com/docker-ce/linux/ubuntu \
    $(lsb_release -cs) \
    stable"
    apt-get update
    apt-get install docker-ce docker-ce-cli containerd.io
    systemctl enable docker
    systemctl start docker
    echo 'docker install successfully'
else
    echo 'docker is installed'
fi

####### 安装docker-compose ######
if !type docker-compose >/dev/null 2>&1
then
    curl -L https://download.fastgit.org/docker/compose/releases/download/1.27.4/docker-compose-`uname -s`-`uname -m` > /usr/local/bin/docker-compose
    chmod +x /usr/local/bin/docker-compose
else
    echo 'docker-compose is installed'
fi

chmod 777 ./app/ccs/dist/ccs
docker-compose up -d