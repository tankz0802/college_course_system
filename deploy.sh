#!/bin/sh
set -e
######### 配置阿里镜像源 ########
sudo cp /etc/apt/sources.list /etc/apt/sources.list.bak
cat>/etc/apt/sources.list<<EOF
deb http://mirrors.aliyun.com/ubuntu/ bionic main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ bionic-security main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ bionic-updates main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ bionic-proposed main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ bionic-backports main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ bionic main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ bionic-security main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ bionic-updates main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ bionic-proposed main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ bionic-backports main restricted universe multiverse
EOF
sudo apt-get update

######### 关闭防火墙 ###########
sudo ufw disable

######### 安装curl ############
sudo apt-get -y install curl

######### 安装docker ###########
if ! type docker >/dev/null 2>&1
then
    sudo apt-get remove docker docker-engine docker.io
    sudo apt-get update
    sudo apt-get -y install \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent \
    software-properties-common
    sudo curl -fsSL https://mirrors.aliyun.com/docker-ce/linux/ubuntu/gpg |  apt-key add -
    add-apt-repository \
    "deb [arch=amd64] https://mirrors.aliyun.com/docker-ce/linux/ubuntu \
    $(lsb_release -cs) \
    stable"
    sudo apt-get update
    sudo apt-get install -y docker-ce docker-ce-cli containerd.io
    sudo systemctl enable docker
    sudo systemctl start docker
    echo 'docker install successfully'
else
    echo 'docker is installed'
fi

####### 安装docker-compose ######
if ! type docker-compose >/dev/null 2>&1
then
    sudo curl -L https://get.daocloud.io/docker/compose/releases/download/1.28.2/docker-compose-`uname -s`-`uname -m` > /usr/local/bin/docker-compose
    sudo chmod +x /usr/local/bin/docker-compose
else
    echo 'docker-compose is installed'
fi

########### 替换IP #############
ip=$(ip -o -4 addr show eth0 | awk '{ split($4, ip_addr, "/"); print ip_addr[1] }')
if test -z "$ip"
then
    ip=$(ip -o -4 addr show ens33 | awk '{ split($4, ip_addr, "/"); print ip_addr[1] }')
    if test -z "$ip"
    then
        echo "网络配置错误"
        exit -1
    fi
fi
sed -i -e "s/[0-9]\{1,3\}\.[0-9]\{1,3\}\.[0-9]\{1,3\}\.[0-9]\{1,3\}/$ip/g" ./app/ccs/proxy.config.json
sed -i -e "s/[0-9]\{1,3\}\.[0-9]\{1,3\}\.[0-9]\{1,3\}\.[0-9]\{1,3\}/$ip/g" ./app/ccs/nginx.conf
sed -i -e "s/[0-9]\{1,3\}\.[0-9]\{1,3\}\.[0-9]\{1,3\}\.[0-9]\{1,3\}/$ip/g" ./server/config/config.ini

touch /etc/docker/daemon.json
echo "{\"registry-mirrors\":[\"https://hub-mirror.c.163.com/\"]}" > /etc/docker/daemon.json
sudo systemctl daemon-reload
sudo systemctl restart docker
sudo docker-compose up -d
sudo docker start ccs-server
echo "部署完成,请访问"$ip":4200进行预览!"
set +e
