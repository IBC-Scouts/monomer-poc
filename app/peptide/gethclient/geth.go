package gethclient

import (
	"context"
	"fmt"
	"strings"

	gn "github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/polymerdao/monomer/app/peptide"

	"github.com/ethereum-optimism/optimism/op-service/client"
)

// NewGethEngineClient creates a new geth EngineClient
//
// This is an implementation of the derive.Engine interface using a geth client.
func NewRPCClient(gethEngineAddr string, gethAuthSecret []byte) (client.RPC, error) {
	if strings.TrimSpace(gethEngineAddr) == "" {
		return nil, fmt.Errorf("geth execution engine address must be non-empty")
	}

	var authSecret [32]byte
	if len(gethAuthSecret) == 0 {
		authSecret = [32]byte{123}
	} else {
		copy(authSecret[:], gethAuthSecret[:min(len(gethAuthSecret), 32)])
	}

	auth := rpc.WithHTTPAuth(gn.NewJWTAuth(authSecret))
	opts := []client.RPCOption{
		client.WithGethRPCOptions(auth),
		client.WithDialBackoff(10),
	}

	logger, _ := peptide.NewCompositeLogger("debug", "module", "gethclient")
	rpcClient, err := client.NewRPC(context.TODO(), logger, gethEngineAddr, opts...)
	if err != nil {
		return nil, err
	}

	return rpcClient, nil
}
