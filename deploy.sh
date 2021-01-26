#!/bin/sh
######### 安装docker ###########
if ! type docker >/dev/null 2>&1
then
    apt-get remove docker docker-engine docker.io
    apt-get update
    curl -fsSL https://mirrors.aliyun.com/docker-ce/linux/ubuntu/gpg |  apt-key add -
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
    curl -L https://get.daocloud.io/docker/compose/releases/download/1.25.4/docker-compose-`uname -s`-`uname -m` > /usr/local/bin/docker-compose
    chmod +x /usr/local/bin/docker-compose
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
local_ip=$(ifconfig | grep '\<inet\>'| grep -v '127.0.0.1' | awk '{ print $2}' | awk 'NR==1')
sed -i "s/ip/$ip/g" ./app/ccs/proxy.config.json
sed -i "s/ip/$ip/g" ./app/ccs/nginx.conf
sed -i "s/ip/$ip/g" ./server/config/config.ini
et ff=unix
ufw disable
touch /etc/docker/daemon.json
echo "{\"registry-mirrors\":[\"https://hub-mirror.c.163.com/\"]}" > /etc/docker/daemon.json
systemctl daemon-reload
systemctl restart docker
docker-compose up -d
docker start ccs-server
echo "部署完成,请访问"$ip":4200进行预览!"
