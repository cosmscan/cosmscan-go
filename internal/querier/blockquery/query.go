package blockquery

import (
	"context"
	"cosmscan-go/internal/client"

	coretypes "github.com/tendermint/tendermint/rpc/core/types"
)

func Block(ctx context.Context, cli *client.Client, height int64) (*coretypes.ResultBlock, error) {
	return cli.RPC.Block(ctx, &height)
}

func LatestBlockNumber(ctx context.Context, cli *client.Client) (int64, error) {
	status, err := cli.RPC.Status(ctx)
	if err != nil {
		return 0, err
	}
	return status.SyncInfo.LatestBlockHeight, nil
}
