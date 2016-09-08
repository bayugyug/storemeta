#!/bin/bash 

apt-get purge -y software-properties-common byobu curl git htop man unzip vim \
python-dev python-pip python-virtualenv python-dev python-pip python-virtualenv \
python2.7 python2.7 libpython2.7-stdlib:amd64 libpython2.7-minimal:amd64 \
libgcc-4.8-dev:amd64 cpp-4.8 libruby1.9.1 perl-modules vim-runtime

apt-get clean autoclean
apt-get autoremove -y

rm -rf /var/lib/{apt,dpkg,cache,log}/
rm -rf /var/{cache,log}
