FROM node:10-alpine as builder
WORKDIR /usr/src/app

COPY . .
RUN npm config set registry http://registry.npm.taobao.org/
RUN cnpm install
RUN cnpm run build:prod

FROM nginx:1.17
COPY --from=builder /usr/src/app/dist/ccs /usr/share/nginx/html
COPY ./nginx.conf /etc/nginx/conf.d/default.conf