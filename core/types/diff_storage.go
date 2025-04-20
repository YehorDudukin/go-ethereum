package types

import (
	"github.com/YehorDudukin/go-ethereum/common"
)

type DiffStorage map[common.Address]*AccountDiff

type AccountDiff struct {
	Data    StateAccount
	Code    []byte
	Storage map[common.Hash]common.Hash
}
