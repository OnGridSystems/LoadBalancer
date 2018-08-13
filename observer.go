package main

import (
	"time"
)

func observeNode(node Node, config Config) Node {
	Info.Printf("Observing node: %s", node.Url.String())
	blockNumber, err := getBlockNumber(&node, config)

	if err != nil {
		Error.Printf("Obsserving failed with: %v", err)
		node.Available = false
	} else {
		node.Available = true
		node.BlockNumber = blockNumber
		Info.Printf("Observing result: %+v", node)
	}

	return node
}

func chooseBestNodeId(nodes []Node, config Config) (bestNodeId int) {
	var maxBlock int64 = 0

	for i, n := range nodes {
		if n.Available && n.BlockNumber > maxBlock {
			maxBlock = n.BlockNumber

			bestNodeId = i
		}

	}

	return bestNodeId
}

func observe(config Config, nodes []Node, currentNodeId *int) {
	for i, node := range nodes {
		nodes[i] = observeNode(node, config)
	}

	bestNodeId := chooseBestNodeId(nodes, config)

	if currentNodeId == nil {
		*currentNodeId = bestNodeId
	} else {
		currentNode := nodes[*currentNodeId]
		bestNode := nodes[bestNodeId]

		blockDiff := currentNode.BlockNumber - bestNode.BlockNumber

		if !currentNode.Available || blockDiff > config.BlockThreshold {
			*currentNodeId = bestNodeId
		}
	}

	Info.Printf("Best node: %+v", nodes[*currentNodeId].Url.String())
}

func startPeriodicObserve(config Config, nodes []Node, currentNodeId *int) {
	ticker := time.NewTicker(time.Duration(config.Interval) * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		observe(config, nodes, currentNodeId)
	}
}
