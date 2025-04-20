package types

import (
	"github.com/YehorDudukin/go-ethereum/common"
)

type DiffStorage []AccountDiff

type AccountDiff struct {
	Address      common.Address
	Account      StateAccount
	Code         []byte
	DirtyStorage []DirtyStorage
}

type DirtyStorage struct {
	Key   common.Hash
	Value common.Hash
}
