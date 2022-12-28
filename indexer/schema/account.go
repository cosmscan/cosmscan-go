package schema

// Accounts is a list of accounts.
type Accounts []Account

// Account represents a single account in the blockchain.
type Account struct {
	Address string
}

// AccountsFromFullBlock returns a list of accounts extracted from the FullBlock
// For now, we only extracted accounts from "transfer" events in a transaction.
func AccountsFromFullBlock(block *FullBlock) Accounts {
	accounts := make(Accounts, 0)

	for _, transaction := range block.Txs {
		for _, evt := range transaction.Events {
			if evt.Type == "transfer" {
				switch evt.Key {
				case "sender":
					accounts = append(accounts, Account{Address: evt.Value})
				case "recipient":
					accounts = append(accounts, Account{Address: evt.Value})
				}
			}
		}
	}

	return accounts
}
