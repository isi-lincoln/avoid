#!/bin/bash

apt-get update
dpkg --configure -a
DEBIAN_FRONTEND=noninteractive apt-get install -qy \
    dnsutils \
    iproute2 \
    iputils-ping \
    jq \
    vim-nox

PREFIX=`ip -j addr | jq '.[] | select(.ifname=="eth0")'.addr_info[0].local | sed 's/"//g' | cut -d "." -f 1-3`

# TODO: this nameserver depends on if coredns gets the first address
cat <<EOF > /etc/resolv.conf
nameserver $PREFIX.2
options ndots:0
EOF

sleep 10d
