#start
FROM golang

#auth
MAINTAINER bayugyug<bayugyug@gmail.com>

#timezone
RUN echo "Asia/Singapore" > /etc/timezone && dpkg-reconfigure -f noninteractive tzdata

# Install apache, PHP, and supplimentary programs. curl and lynx-cur are for debugging the container.
RUN apt-get update -y && apt-get install -y --no-install-recommends \
        curl vim \
    && rm -rf /var/lib/apt/lists/*

RUN mkdir -p /etc/storemeta

ADD storemeta /bin/storemeta
RUN chmod +x /bin/storemeta

ENTRYPOINT ["storemeta"]
