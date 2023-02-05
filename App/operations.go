package App

import (
	"fmt"

	"github.com/LukaPetkovicSV16/Projekat-NAISP/lru"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/memtable"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/sstable"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/types"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/wal"
)

func HandleAdd(key string, value []byte, memtable *memtable.Memtable, LRU *lru.LRUCache) {
	// TODO: check if request can be made with token bucket
	var newRecord = types.CreateRecord(key, value, false)
	if wal.Append(newRecord) {
		memtable.Add(newRecord)
		LRU.Add(newRecord)
	}
}
func HandleGet(key string, memtable *memtable.Memtable, LRU *lru.LRUCache) *types.Record {
	// TODO: check if request can be made with token bucket
	var record = memtable.Get(key)
	if record != nil {
		if record.Tombstone {
			return nil
		}
		return record
	}
	var checkCache = LRU.Read(key)
	if checkCache != nil {
		return checkCache
	}
	return sstable.Read(key)
}
func HandleDelete(key string, memtable *memtable.Memtable, LRU *lru.LRUCache) bool {
	// TODO: check if request can be made with token bucket
	var record = HandleGet(key, memtable, LRU)
	if record != nil {
		if !record.Tombstone {
			record.Tombstone = true
			if wal.Append(*record) {
				memtable.Add(*record)
				return true
			}
		}
	} else {
		fmt.Println("Record doesn't exist or is deleted")
	}
	return false
}

func HandleGetList() {
	// TODO: Implement this
}
func HandleRangeScan() {
	// TODO: Implement this
}
