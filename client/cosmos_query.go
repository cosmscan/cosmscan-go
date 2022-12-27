package client

import (
	"context"

	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
)

func TransactionByHash(ctx context.Context, hash string, cli *Client) (*txtypes.GetTxResponse, error) {
	request := &txtypes.GetTxRequest{Hash: hash}
	svc := txtypes.NewServiceClient(cli)
	return svc.GetTx(ctx, request)
}
