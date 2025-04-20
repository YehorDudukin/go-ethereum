package types

import (
	"github.com/YehorDudukin/go-ethereum/common"
)

type DiffStorage map[common.Address]map[common.Hash]common.Hash

func (d DiffStorage) DeepCopy() DiffStorage {
	copyStorage := make(DiffStorage, len(d))
	for addr, slots := range d {
		slotCopy := make(map[common.Hash]common.Hash, len(slots))
		for k, v := range slots {
			slotCopy[k] = v
		}
		copyStorage[addr] = slotCopy
	}

	return copyStorage
}
