package fetcher

import (
	"context"
	"cosmscan-go/client"
	"cosmscan-go/indexer/schema"
	"errors"
	"sync"

	"go.uber.org/zap"
)

// AccountBalanceFetcher aims to keep fetching the balances with accounts
// we want to track how many account exists in the blockchain or who is top 10 holders.
type AccountBalanceFetcher struct {
	cli        *client.Client
	log        *zap.SugaredLogger
	ctx        context.Context
	cancelFunc context.CancelFunc
	init       bool
	subOnce    sync.Once
	reqC       chan *schema.Account
	resC       chan *schema.AccountBalance
}

func NewAccountBalanceFetcher(cli *client.Client) *AccountBalanceFetcher {
	ctx, cancel := context.WithCancel(context.Background())

	return &AccountBalanceFetcher{
		cli:        cli,
		log:        zap.S().Named("account_balance_fetcher"),
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
			f.log.Info("the system is about to stop")
			return
		case account := <-f.reqC:
			res, err := client.BalanceOf(f.ctx, f.cli, account.Address)
			if err != nil {
				f.log.Warn("failed to get balance of account and skipped fetching for this acc", "account", account.Address, "err", err)
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
