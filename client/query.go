package client

import (
	"context"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
)

func (c *Client) Block(ctx context.Context, height int64) (*coretypes.ResultBlock, error) {
	return c.rpc.Block(ctx, &height)
}

func (c *Client) Status(ctx context.Context) (*coretypes.ResultStatus, error) {
	return c.rpc.Status(ctx)
}
