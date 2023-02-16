package fetcher

import (
	"context"
	client2 "cosmscan-go/internal/client"
	"cosmscan-go/modules/indexer/schema"
	"cosmscan-go/pkg/log"
	"errors"
	"sync"
)

// AccountBalanceFetcher aims to keep fetching the balances with accounts
// we want to track how many account exists in the blockchain or who is top 10 holders.
type AccountBalanceFetcher struct {
	cli        *client2.Client
	ctx        context.Context
	cancelFunc context.CancelFunc
	init       bool
	subOnce    sync.Once
	reqC       chan *schema.Account
	resC       chan *schema.AccountBalance
}

func NewAccountBalanceFetcher(cli *client2.Client) *AccountBalanceFetcher {
	ctx, cancel := context.WithCancel(context.Background())

	return &AccountBalanceFetcher{
		cli:        cli,
		ctx:        ctx,
		cancelFunc: cancel,
		init:       false,
	}
}

func (f *AccountBalanceFetcher) Subscribe() (requestCh chan<- *schema.Account, responseCh <-chan *schema.AccountBalance, err error) {
	if f.init {
		return nil, nil, errors.New("already subscribed")
	}

	f.subOnce.Do(func() {
		f.reqC = make(chan *schema.Account)
		f.resC = make(chan *schema.AccountBalance)
		f.init = true
	})

	return f.reqC, f.resC, nil
}

func (f *AccountBalanceFetcher) Close() {
	f.cancelFunc()
}

func (f *AccountBalanceFetcher) Run() {
	f.run()
}

func (f *AccountBalanceFetcher) run() {
	for {
		select {
		case <-f.ctx.Done():
			log.Logger.Info("the system is about to stop")
			return
		case account := <-f.reqC:
			res, err := client2.BalanceOf(f.ctx, f.cli, account.Address)
			if err != nil {
				log.Logger.Warn("failed to get balance of account and skipped fetching for this acc", "account", account.Address, "err", err)
				continue
			}

			var balances schema.AccountCoins
			for _, b := range res.Balances {
				balances = append(balances, &schema.AccountCoin{
					Denom:  b.Denom,
					Amount: b.Amount.BigInt(),
				})
			}

			f.resC <- &schema.AccountBalance{
				Account: account,
				Balance: balances,
			}
		}
	}
}
