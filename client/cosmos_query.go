package client

import (
	"context"

	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// TransactionByHash returns a transaction by hash
func TransactionByHash(ctx context.Context, hash string, cli *Client) (*txtypes.GetTxResponse, error) {
	request := &txtypes.GetTxRequest{Hash: hash}
	svc := txtypes.NewServiceClient(cli)
	return svc.GetTx(ctx, request)
}

func BalanceOf(ctx context.Context, cli *Client, address string) (*banktypes.QueryAllBalancesResponse, error) {
	request := &banktypes.QueryAllBalancesRequest{
		Address:    address,
		Pagination: nil,
	}
	svc := banktypes.NewQueryClient(cli)
	res, err := svc.AllBalances(ctx, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}
