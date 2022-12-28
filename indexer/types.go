package indexer

import (
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
)

type CommitBlockChannel chan *msgCommitBlock

type rawTx struct {
	abci   *coretypes.ResultTx
	cosmos *txtypes.GetTxResponse
}

type msgCommitBlock struct {
	block *coretypes.ResultBlock
	txs   []*rawTx
}
