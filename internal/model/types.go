package model

type TxType int
type BlockHeight uint64

const (
	NormalTx TxType = iota
	BeginBlock
	EndBlock
)
