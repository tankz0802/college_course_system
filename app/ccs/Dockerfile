# 暂时不用该Dockerfile构建,因为在云服务器上对项目进行编译打包会把内存撑爆
FROM node:12-alpine as builder
WORKDIR /usr/src/app

COPY . .
RUN npm config set registry http://registry.npm.taobao.org/
RUN cnpm install -g @angular/cli@9.0.7
RUN cnpm install
RUN cnpm run build:prod

FROM nginx:1.17
COPY --from=builder /usr/src/app/dist/ccs /usr/share/nginx/html
COPY ./nginx.conf /etc/nginx/conf.d/default.conf