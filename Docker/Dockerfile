#start
FROM golang

#auth
MAINTAINER bayugyug<bayugyug@gmail.com>

# Golang
RUN apt-get update -y && apt-get install -y --no-install-recommends \
    git \
    && rm -rf /var/lib/apt/lists/*

#from remote
RUN mkdir -p /go/src/github.com/bayugyug

RUN cd /go/src/github.com/bayugyug && git clone https://github.com/bayugyug/storemeta.git

#compile
RUN cd /go/src/github.com/bayugyug/storemeta && go get -v && go build -v -ldflags "-X main.pBuildTime=`date -u +%Y%m%d.%H%M%S`"

RUN cd /go/src && go install github.com/bayugyug/storemeta

RUN chmod +x /go/bin/storemeta

#free
RUN apt-get clean

RUN apt-get purge

ENTRYPOINT ["/go/bin/storemeta"]