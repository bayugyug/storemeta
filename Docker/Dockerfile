#start
FROM golang

#auth
MAINTAINER bayugyug<bayugyug@gmail.com>

#prepare
RUN apt-get update -y && apt-get install -y --no-install-recommends \
git \
&& rm -rf /var/lib/apt/lists/*

#from remote
RUN mkdir -p /go/src/github.com/bayugyug

RUN cd /go/src/github.com/bayugyug && git clone https://github.com/bayugyug/storemeta.git

#compile
RUN cd /go/src/github.com/bayugyug/storemeta && go get -v && go build -v -ldflags "-X main.pBuildTime=`date -u +%Y%m%d.%H%M%S`"

#install
RUN cd /go/src && go install github.com/bayugyug/storemeta

#ADD
RUN cd /go/src/github.com/bayugyug/storemeta  && cp -f Docker/cleanup.sh /go/bin/

#executable
RUN chmod +x /go/bin/*

#cleanup
RUN /go/bin/cleanup.sh

#main entry
ENTRYPOINT ["/go/bin/storemeta"]
