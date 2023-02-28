package schema

import (
	"cosmscan-go/internal/model"
	"fmt"
	"time"

	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
)

type FullBlock struct {
	Block *model.Block
	Txs   []*FullTransaction
}

type FullTransaction struct {
	Tx       *model.Transaction
	Events   []*model.Event
	Messages []*model.Message
}

func NewFullBlock(block *coretypes.ResultBlock, abciTx []*coretypes.ResultTx, cosmTx []*txtypes.GetTxResponse) (*FullBlock, error) {
	txs := make([]*FullTransaction, 0)

	// fill block
	b := &model.Block{
		Height:             model.BlockHeight(block.Block.Height),
		Hash:               block.Block.Hash().String(),
		ProposerAddress:    block.Block.ProposerAddress.String(),
		LastCommitHash:     block.Block.LastCommitHash.String(),
		DataHash:           block.Block.DataHash.String(),
		ValidatorsHash:     block.Block.ValidatorsHash.String(),
		NextValidatorsHash: block.Block.NextValidatorsHash.String(),
		ConsensusHash:      block.Block.ConsensusHash.String(),
		AppHash:            block.Block.AppHash.String(),
		LastResultHash:     block.Block.LastResultsHash.String(),
		EvidenceHash:       block.Block.EvidenceHash.String(),
		BlockTime:          block.Block.Time,
	}

	for i, tx := range abciTx {
		fullTx := &FullTransaction{
			Events:   make([]*model.Event, 0),
			Messages: make([]*model.Message, 0),
		}
		res := cosmTx[i]

		txTime, err := time.Parse(time.RFC3339, res.TxResponse.Timestamp)
		if err != nil {
			return nil, err
		}

		// fill transaction
		fullTx.Tx = &model.Transaction{
			Hash:      tx.Hash.String(),
			Height:    model.BlockHeight(block.Block.Height),
			Seq:       int(tx.Index),
			Code:      int(res.TxResponse.Code),
			CodeSpace: res.TxResponse.Codespace,
			TxData:    res.TxResponse.Data,
			RawLog:    res.TxResponse.RawLog,
			Info:      res.TxResponse.Info,
			Memo:      res.Tx.Body.Memo,
			GasWanted: uint64(res.TxResponse.GasWanted),
			GasUsed:   uint64(res.TxResponse.GasUsed),
			IssuedAt:  txTime,
		}

		// fill events
		for _, txLog := range res.TxResponse.Logs {
			for _, event := range txLog.Events {
				for _, attr := range event.Attributes {
					fullTx.Events = append(fullTx.Events, &model.Event{
						Height:  model.BlockHeight(block.Block.Height),
						Seq:     txLog.MsgIndex,
						Type:    event.Type,
						Key:     attr.Key,
						Value:   attr.Value,
						Indexed: false,
					})
				}
			}
		}

		// fill messages
		for seq, msg := range res.Tx.Body.Messages {
			// TODO: wanna to store the message as raw json format
			// e.g. { "type": "MsgSend", "value": { "sender": "alice", "recipient": "bob", "amount": "1000" } }
			fullTx.Messages = append(fullTx.Messages, &model.Message{
				Seq:     seq,
				RawData: fmt.Sprintf("{ \"type\": \"%s\"}", msg.TypeUrl),
			})
		}

		txs = append(txs, fullTx)
	}

	return &FullBlock{
		Block: b,
		Txs:   txs,
	}, nil
}
