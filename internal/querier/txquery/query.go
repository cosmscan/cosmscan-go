package txquery

import (
	"context"
	"cosmscan-go/internal/client"

	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
)

// TransactionByHash returns a transaction by hash
func TransactionByHash(ctx context.Context, cli *client.Client, hash string) (*txtypes.GetTxResponse, error) {
	request := &txtypes.GetTxRequest{Hash: hash}
	svc := txtypes.NewServiceClient(cli)
	return svc.GetTx(ctx, request)
}

func ABCITransactionByHash(ctx context.Context, cli *client.Client, hash []byte) (*coretypes.ResultTx, error) {
	return cli.RPC.Tx(ctx, hash, false)
}
