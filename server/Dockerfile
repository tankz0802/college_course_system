FROM golang:alpine as build

ENV GOPROXY https://goproxy.io
ENV GO111MODULE on

WORKDIR /go/cache

ADD go.mod .
ADD go.sum .
RUN go mod download

WORKDIR /go/release

ADD . .

RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o ccs main.go

FROM ubuntu as prod

COPY --from=build /go/release/ccs /
COPY --from=build /go/release/config /config
COPY --from=build /go/release/assets /assets
CMD ["/ccs"]
