FROM golang:1.18

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/per
COPY . $GOPATH/src/per
RUN go build .
EXPOSE 55555
ENTRYPOINT ["./permissions"]