package accountquery

import (
	"context"
	"cosmscan-go/internal/client"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func BalanceOf(ctx context.Context, cli *client.Client, address string) (*banktypes.QueryAllBalancesResponse, error) {
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
