#start
#FROM golang
FROM google/debian:wheezy

#auth
MAINTAINER bayugyug<bayugyug@gmail.com>

# Golang
RUN apt-get update -y && apt-get install -y --no-install-recommends \
    git \
    && rm -rf /var/lib/apt/lists/*

RUN mkdir -p /go/bin /go/src

ADD storemeta /go/bin/storemeta

RUN chmod +x /go/bin/storemeta

#free
RUN apt-get clean

RUN apt-get purge

ENTRYPOINT ["/go/bin/storemeta"]
