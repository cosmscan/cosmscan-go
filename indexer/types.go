package indexer

import (
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
)

type tx struct {
	abci   *coretypes.ResultTx
	cosmos *txtypes.GetTxResponse
}

type preCommitBlock struct {
	block *coretypes.ResultBlock
	txs   []*tx
}
