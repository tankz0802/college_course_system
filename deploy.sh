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

touch /etc/docker/daemon.json
echo "{\"registry-mirrors\":[\"https://mirror.ccs.tencentyun.com\"]}" > /etc/docker/daemon.json
sudo systemctl daemon-reload
sudo systemctl restart docker
sudo docker build -t ccs .
sudo docker run -d --privileged=true --name=ccs -p 4200:4200 ccs
sudo docker exec -it ccs /bin/sh
docker-compose up
exit
echo "部署完成,请访问127.0.0.1:4200进行预览!"
set +e
