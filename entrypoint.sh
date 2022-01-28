#!/bin/sh

kpng kube --service-proxy-name=$PROXY_NAME to-api &

exec /main
