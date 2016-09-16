#start
#FROM golang
FROM golang:alpine

#auth
MAINTAINER bayugyug<bayugyug@gmail.com>

RUN apk add --no-cache ca-certificates

RUN set -ex \
    && apk add --no-cache --virtual .build-deps \
        bash \
        git \
        gcc \
        musl-dev \
        openssl \
        go

RUN apk add --update ca-certificates # Certificates for SSL

ENV GOPATH /go
ENV GOSRC $GOPATH/src
ENV GOBIN $GOPATH/bin
ENV PATH $PATH:$GOPATH/bin:/usr/local/go/bin

#prepare
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

#home
RUN mkdir -p $GOSRC/github.com/bayugyug && cd $GOSRC/github.com/bayugyug && git clone https://github.com/bayugyug/storemeta.git

#compile
RUN cd $GOSRC/github.com/bayugyug/storemeta && go get -v && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -v -ldflags "-w -X main.pBuildTime=`date -u +%Y%m%d.%H%M%S`" .

#ADD
RUN cd $GOSRC/github.com/bayugyug/storemeta && cp -f storemeta $GOBIN/

#executable
RUN chmod +x $GOBIN/*

#cleanup
RUN du -sh /go/* /* && ls -ltrha $GOBIN/*

RUN apk del .build-deps go git && \
    rm -rf /var/lib/{apt,dpkg,cache,log}/ &&  \
    rm -rf /var/{cache,log} && \
    rm -fr /go/pkg  /go/src

#cleanup
RUN du -sh /go/* /*

#main entry
ENTRYPOINT [ "/go/bin/storemeta" ]
