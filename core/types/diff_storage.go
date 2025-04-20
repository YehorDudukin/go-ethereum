package types

import (
	"github.com/YehorDudukin/go-ethereum/common"
)

type DiffStorage map[common.Address]map[common.Hash]common.Hash
