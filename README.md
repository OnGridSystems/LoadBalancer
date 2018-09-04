[![OnGrid Systems Blockchain Applications DApps Development](img/ongrid-systems-cover.png)](https://ongrid.pro/)

# Load Balancer
A simple load balancer / reverse proxy written in GoLang for blockchain nodes.
Only Web3 compatible nodes are currently supported.

## Features
 - Acts as a reverse proxy, sitting between clients and web servers.
 - Chooses node with biggest block number
 - Falls back on other web-servers if one goes down

# Installation
With docker

```
docker build . -t loadbalancer
```

Without docker
```
go get -u github.com/golang/dep/cmd/dep
dep ensure
go build
```

# Using
Example config file
```
port: 8000
check_interval: 30  # seconds
connection_timeout: 5 # seconds
nodes:
  - https://mainnet.infura.io/token
  - http://localhost:8545
block_treshold: 10
```
* port - listening port
* check_interval - nodes polling interval
* connection_timeout - nodes polling connection timeout
* nodes - list of polling nodes
* block_treshold - node switch block treshold

Start balancer:
With docker
```
docker run -v config.yml:/config.yml loadbalancer
```
Without docker
```
LoadBalancer -c /path/to/config.yml
```

## Future Features
[Issues tagged 'feature'](https://github.com/nwoodthorpe/Load-Balancer-Golang/issues?q=is%3Aissue+is%3Aopen+label%3Afeature)
