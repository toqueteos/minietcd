#!/bin/sh

set -e

ETCD_CACHE_DIR=${HOME}/etcd

if [ ! -d ${ETCD_CACHE_DIR} ]; then
    mkdir ${ETCD_CACHE_DIR};
    cd ${ETCD_CACHE_DIR};
    curl -L https://github.com/coreos/etcd/releases/download/v2.2.5/etcd-v2.2.5-linux-amd64.tar.gz -o ${ETCD_CACHE_DIR}/etcd-v2.2.5-linux-amd64.tar.gz
    tar xzvf etcd-v2.2.5-linux-amd64.tar.gz;
else
    echo 'Using cached etcd';
fi
