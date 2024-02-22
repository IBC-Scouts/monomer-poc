// Poc rpc holds some rpc methods for interacting with peptide via rpc (mainly, adding to its mempool
// a-la eth_sendRawTransaction).
package poc_rpc

import (
	bfttypes "github.com/cometbft/cometbft/types"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/polymerdao/monomer/app/node/server"
)

// MempoolNode is an interface that should be implemented by the node to allow the rpc to add
// transactions to the mempool.
type MempoolNode interface {
	// AddToTxMempool adds a transaction to the mempool.
	AddToTxMempool(bfttypes.Tx)
}

// GetInterceptorEntrypointAPI returns the rpc API for the interceptor entrypoints.
func GetInterceptorEntrypointAPI(node MempoolNode, logger server.Logger) rpc.API {
	return rpc.API{
		Namespace: "intercept",
		Service:   &interceptEndpoints{node: node, logger: logger.With("module", "intercept")},
	}
}

// interceptEndpoints is the struct that holds the rpc methods for the interceptor entrypoints.
// These are methods the interceptor can invoke in order to manipulate the node's state.
// (Note: we shoud be able to use ABCI for this too but we shouldn't really need the additional methods plus
// we already have an rpc client for the node so we might as well use it.)
type interceptEndpoints struct {
	node   MempoolNode
	logger server.Logger
}

// AddToTxMempool adds a transaction to the mempool. It takes raw bytes containing the marshalled
// sdk.Msg and should try to create a Tx from it.
func (p *interceptEndpoints) AddMsgToTxMempool(msg []byte) error {
	p.logger.Info("Adding msg to mempool", "msg", msg)

	// p.node.AddToTxMempool(tx)

	return nil
}
