package App

import (
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
		// LRU.Add(newRecord)
	}
}
func HandleGet(key string, memtable *memtable.Memtable, LRU *lru.LRUCache) *types.Record {
	// TODO: check if request can be made with token bucket
	var record = memtable.Get(key)
	if record != nil {
		return record
	}
	// var checkCache = LRU.Read(key)
	// if checkCache != nil {
	// 	return checkCache
	// }
	return sstable.Read(key)
}
func HandleDelete(key string, memtable *memtable.Memtable) bool {
	// TODO: check if request can be made with token bucket
	//TODO: check if key exist in memtable and set tombstone to true
	// var record = HandleGet(key, memtable)
	// if record != nil {
	// 	record.Tombstone = true
	// if wal.Append(newRecord) {
	// 	memtable.Add(key, newRecord)
	// }
	// }else{
	// fmt.Println("Record doesn't exist")
	// }
	return false
}

func HandleGetList() {
	// TODO: Implement this
}
func HandleRangeScan() {
	// TODO: Implement this
}
