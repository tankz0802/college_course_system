FROM node:12-alpine as builder
WORKDIR /usr/src/app

COPY . .
RUN npm install -g cnpm --registry http://registry.npm.taobao.org/
RUN cnpm install -g @angular/cli@9.0.7
RUN cnpm install
RUN ng build --prod

FROM nginx:1.17
COPY --from=builder /usr/src/app/dist/ccs /usr/share/nginx/html
COPY ./nginx.conf /etc/nginx/conf.d/default.conf