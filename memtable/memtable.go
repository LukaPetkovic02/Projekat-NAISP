package memtable

import "github.com/LukaPetkovicSV16/Projekat-NAISP/types"

// TODO: implement all commented methods in skipList and bTree
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
	// TODO: Check if memtable is full if it is Flush it here
	return memtable.Records.Add(key, record)
}

func (memtable *Memtable) Delete(key string) bool {
	return true
}

func Flush() {
	// Call write to ssTable here
}
