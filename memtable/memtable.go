package memtable

import (
	"github.com/LukaPetkovicSV16/Projekat-NAISP/engine"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/sstable"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/types"
)

type Data interface {
	Get(key string) *types.Record
	Add(key string, record types.Record) bool
	Delete(key string) bool
	GetSortedRecordsList() []types.Record
	Clear()
	GetSize() int
}

type Memtable struct {
	MaxSize int // Max size of memtable in bytes
	Records Data
}

func Init(maxSize int, records Data) *Memtable {
	return &Memtable{
		MaxSize: maxSize,
		Records: records,
	}
}

func (memtable *Memtable) Get(key string) *types.Record {
	return memtable.Records.Get(key)
}

func (memtable *Memtable) Add(key string, record types.Record) bool {
	if memtable.MaxSize <= memtable.Records.GetSize()+engine.DEFAULT_MEMTABLE_THRESHOLD {
		memtable.Flush()
	}
	return memtable.Records.Add(key, record)
}

func (memtable *Memtable) Delete(key string) bool {
	return memtable.Records.Delete(key)
}

func (memtable *Memtable) Flush() {
	var records = memtable.Records.GetSortedRecordsList()
	sstable.Create(records)
	memtable.Records.Clear()
}
