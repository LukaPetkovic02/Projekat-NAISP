package memtable

import "github.com/LukaPetkovicSV16/Projekat-NAISP/types"

type Data interface {
	Get(key string) *types.Record
	Add(key string, record types.Record) bool
	// Delete(key string) bool
	// GetSortedRecordsList() []types.Record
	// Clear()
	// GetSize() int
}

type Memtable struct {
	RecordCapacity int
	Records        Data
}

func Init(recordCapacity int, records Data) *Memtable {
	return &Memtable{
		RecordCapacity: recordCapacity,
		Records:        records,
	}
}

func (memtable *Memtable) Get(key string) *types.Record {
	return memtable.Records.Get(key)
}

func (memtable *Memtable) Add(key string, record types.Record) bool {
	return memtable.Records.Add(key, record)
}
