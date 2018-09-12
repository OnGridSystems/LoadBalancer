[![OnGrid Systems Blockchain Applications DApps Development](img/ongrid-systems-cover.png)](https://ongrid.pro/)

![GitHub last commit](https://img.shields.io/github/last-commit/OnGridSystems/LoadBalancer.svg)
[![Ethereum grant status](https://img.shields.io/badge/language-golang-green.svg)](https://blog.ethereum.org/2018/08/17/ethereum-foundation-grants-update-wave-3/)
![License](https://img.shields.io/github/license/OnGridSystems/LoadBalancer.svg)

# Blockchain Load Balancer
Load balancer / reverse proxy written in GoLang for blockchain nodes.
It provides redundant interface for applications using Ethereum and acts as a reverse proxy 
sitting between ethereum client and nodes. It constantly checks the node availability and its latest block number and
keeps the list of healthy web3 providers. If node goes offline or slows down the requests fall back to another node.
Only Ethereum Web3 over HTTPs nodes are currently supported (see [roadmap](https://github.com/nwoodthorpe/Load-Balancer-Golang/issues?q=is%3Aissue+is%3Aopen+label%3Afeature))

## Install
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

## Configure
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

## Run 
With docker
```
docker run -v config.yml:/config.yml loadbalancer
```
Without docker
```
LoadBalancer -c /path/to/config.yml
```

### License

Each file included in this repository is licensed under the [MIT license](LICENSE).