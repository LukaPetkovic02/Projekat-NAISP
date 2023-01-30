package App

import (
	"github.com/LukaPetkovicSV16/Projekat-NAISP/memtable"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/types"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/wal"
)

func HandleAdd(key string, value []byte, memtable *memtable.Memtable) {
	var newRecord = types.CreateRecord(key, value, false)
	if wal.Append(newRecord) {
		memtable.Add(key, newRecord)
	}
}
func HandleGet(key string, memtable *memtable.Memtable) *types.Record {
	var record = memtable.Get(key)
	if record != nil {
		return record
	}
	// TODO: check cache
	// TODO: check disk
	return nil
}
func HandleDelete(key string, memtable *memtable.Memtable) {
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
}