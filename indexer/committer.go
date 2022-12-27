package indexer

import "cosmscan-go/db"

type Committer struct {
	storage db.DB
	blockCh chan *preCommitBlock
}

func NewCommitter(storage db.DB) *Committer {
	return &Committer{
		storage: storage,
	}
}

func (c *Committer) BlockCh() <-chan *preCommitBlock {
	return c.blockCh
}
