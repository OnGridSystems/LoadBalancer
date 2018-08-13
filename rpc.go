package main

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
	"time"
)

type JSONRPCRequest struct {
	Version string   `json:"jsonrpc"`
	Method  string   `json:"method"`
	Params  []string `json:"params"`
	Id      int      `json:"id"`
}

type JSONRPCResponse struct {
	Version string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Result  string `json:"result"`
}

func getBlockNumber(node *Node, config Config) (int64, error) {
	request := JSONRPCRequest{
		Version: "2.0",
		Method:  "eth_blockNumber",
		Id:      node.RPCCounter,
		Params:  make([]string, 0),
	}

	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(request)
	if err != nil {
		return 0, err
	}

	client := &http.Client{
		Timeout: time.Duration(config.Interval) * time.Second,
	}
	resp, err := client.Post(node.Url.String(), "application/json", body)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return 0, errors.Errorf("Invalid response status: %d", resp.StatusCode)
	}
	node.RPCCounter += 1

	var response JSONRPCResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return 0, err
	}

	return strconv.ParseInt(response.Result, 0, 64)
}
