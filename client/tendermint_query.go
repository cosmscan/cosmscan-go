package client

import (
	"context"

	coretypes "github.com/tendermint/tendermint/rpc/core/types"
)

func (c *Client) Block(ctx context.Context, height int64) (*coretypes.ResultBlock, error) {
	return c.rpc.Block(ctx, &height)
}

func (c *Client) ABCITransactionByHash(ctx context.Context, hash []byte) (*coretypes.ResultTx, error) {
	return c.rpc.Tx(ctx, hash, false)
}

func (c *Client) LatestBlockNumber(ctx context.Context) (int64, error) {
	status, err := c.rpc.Status(ctx)
	if err != nil {
		return 0, err
	}
	return status.SyncInfo.LatestBlockHeight, nil
}
