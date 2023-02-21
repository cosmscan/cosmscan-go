package model

type TxType int
type BlockHeight uint64

const (
	NormalTx TxType = iota
	BeginBlock
	EndBlock
)

// ModelsToAutoMigrate returns auto-migration target models
func ModelsToAutoMigrate() []interface{} {
	return []interface{}{
		&Account{},
		&AccountBalance{},
		&Block{},
		&Chain{},
		&Event{},
		&Message{},
		&Transaction{},
	}
}
