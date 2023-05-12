#!/bin/bash

#build
make -C ../ build

# # consul instance
# # -v /Users/<your_name>/tmp/consul:/consul/data \
docker run -d --name dev-consul -p 8500:8500 -p 8600:8600/udp \
-e CONSUL_BIND_INTERFACE=eth0 \
consul agent --bootstrap -server -ui -client=0.0.0.0

# Set consul configuration
root="go-micro-layouts"
appname="grpc-layout"
# Consul url
url=http://host.docker.internal:8500/v1/kv/${root}
token=$1
requestbody='{
  "server": {
    "name": "grpc-layout",
    "addr": ":9090"
  },
  "metric": {
    "addr": ":8080"
  }
}'

# Configuration
# --- DO NOT EDIT BELOW ---
setConsulConfig () {
  echo "### Setting ${root}/${appname} as:"
  echo "${requestbody}"
  if [[ "$(curl -sX PUT -H "X-Consul-Token: ${token}" -d "${requestbody}" ${url}/${appname})" == "true" ]]; then
    echo "### ${url}/${appname} is set"
  else
    echo "### ERROR: Cannot set ${url}/${appname}"
    exit 1
  fi
}
setConsulConfig

#run service
../bin/app
