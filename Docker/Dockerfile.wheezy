#start
#FROM golang
FROM google/debian:jessie

#auth
MAINTAINER bayugyug<bayugyug@gmail.com>

#prepare
RUN apt-get update -y && apt-get install -y --no-install-recommends \
curl tar git ca-certificates apt-transport-https \
&& rm -rf /var/lib/apt/lists/*

#from remote
RUN curl --insecure -O https://storage.googleapis.com/golang/go1.6.linux-amd64.tar.gz
RUN tar -xvf go1.6.linux-amd64.tar.gz
RUN mv -f go /usr/local && rm -f go1.6.linux-amd64.tar.gz


ENV GOHOME=/mongers
ENV GOROOT=/usr/local/go
ENV PATH=$PATH:$GOROOT/bin
ENV GOPATH=$GOHOME/go
ENV GOSRC=$GOPATH/src


RUN mkdir -p $GOHOME/go/src
RUN mkdir -p $GOPATH/src/github.com/bayugyug

#copy
COPY cleanup.sh $GOPATH/bin/

RUN cd $GOSRC/github.com/bayugyug && git clone https://github.com/bayugyug/storemeta.git

#compile
RUN cd $GOSRC/github.com/bayugyug/storemeta && go get -v && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -tags netgo -installsuffix netgo -v -ldflags "-X main.pBuildTime=`date -u +%Y%m%d.%H%M%S`" .

#install
RUN cd $GOSRC && go install github.com/bayugyug/storemeta

#ADD
RUN cd $GOSRC/github.com/bayugyug/storemeta  && cp -f storemeta Docker/cleanup.sh $GOPATH/bin/

#executable
RUN chmod +x $GOPATH/bin/* && rm -fr $GOSRC

#cleanup
RUN $GOPATH/bin/cleanup.sh && rm -fr $GOSRC

#main entry
ENTRYPOINT [ "/mongers/go/bin/storemeta" ]
