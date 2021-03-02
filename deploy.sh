#!/bin/sh
set -e
######### 安装docker ###########
if ! type docker >/dev/null 2>&1
then
    curl -fsSL get.docker.com -o get-docker.sh
    sudo sh get-docker.sh --mirror Aliyun
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
sudo sleep 5
sudo docker exec -it ccs /bin/sh -c "docker-compose up -d"
echo "部署完成,请访问127.0.0.1:4200进行预览!"
set +e
