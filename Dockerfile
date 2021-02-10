FROM docker:dind
RUN touch /etc/docker/daemon.json
RUN echo "{\"registry-mirrors\":[\"https://hub-mirror.c.163.com/\"]}" > /etc/docker/daemon.json
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add curl
RUN curl -L https://get.daocloud.io/docker/compose/releases/download/1.28.2/docker-compose-`uname -s`-`uname -m` > /usr/local/bin/docker-compose
RUN chmod +x /usr/local/bin/docker-compose
WORKDIR /college_course_system
ADD . .
EXPOSE 4200
CMD ["docker-compose up -d"]