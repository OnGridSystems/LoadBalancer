package main

import (
	"flag"
	"net/url"
	"os"
)

type Node struct {
	Url         url.URL
	BlockNumber int64
	Available   bool
	RPCCounter  int
}

func initNodes(config Config) []Node {
	nodes := make([]Node, len(config.Nodes))

	for i, n := range config.Nodes {
		if url, err := url.Parse(n); err == nil {
			nodes[i] = Node{
				Url:         *url,
				BlockNumber: 0,
				Available:   false,
				RPCCounter:  0,
			}
		} else {
			panic(err)
		}
	}

	return nodes
}

func main() {
	InitLogger(os.Stdout, os.Stdout, os.Stderr)

	configPath := flag.String("config", "config.yml", "Path to configuration file")
	flag.Parse()

	config := ParseConfigWPanic(*configPath)
	Info.Printf("Config: %+v\n", config)

	nodes := initNodes(config)

	var currentNodeId int

	observe(config, nodes, &currentNodeId)
	go startPeriodicObserve(config, nodes, &currentNodeId)

	startProxy(config.Port, nodes, &currentNodeId)
}
