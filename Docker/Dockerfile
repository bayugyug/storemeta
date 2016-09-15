#start
FROM golang

#auth
MAINTAINER bayugyug<bayugyug@gmail.com>

#prepare
RUN apt-get update -y && apt-get install -y --no-install-recommends \
git \
&& rm -rf /var/lib/apt/lists/*

#environment 
ENV GOHOME=
ENV GOPATH=$GOHOME/go
ENV GOSRC=$GOPATH/src
ENV PATH=$PATH:$GOPATH/bin

RUN mkdir -p $GOPATH/bin  $GOSRC/github.com/bayugyug

RUN cd $GOSRC/github.com/bayugyug && git clone https://github.com/bayugyug/storemeta.git

#compile
RUN cd $GOSRC/github.com/bayugyug/storemeta && go get -v && go build -tags netgo -installsuffix netgo -installsuffix cgo -v -ldflags "-X main.pBuildTime=`date -u +%Y%m%d.%H%M%S`"

#install
RUN cd $GOSRC && go install github.com/bayugyug/storemeta

#ADD
RUN cd $GOSRC/github.com/bayugyug/storemeta  && cp -f Docker/cleanup.sh $GOPATH/bin/

#executable
RUN chmod +x $GOPATH/bin/*

#cleanup
RUN $GOPATH/bin/cleanup.sh

#main entry
ENTRYPOINT ["storemeta"]
