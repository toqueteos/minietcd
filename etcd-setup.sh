#!/bin/sh

set -e

cd $HOME/etcd;

./etcd &;

./etcdctl --timeout "5s" mkdir foo
./etcdctl --timeout "5s" set foo/foo "Hello"
./etcdctl --timeout "5s" set foo/bar 2
./etcdctl --timeout "5s" set foo/qux 4.5
./etcdctl --timeout "5s" mkdir bar/qux
./etcdctl --timeout "5s" set bar/qux/hello "World"
