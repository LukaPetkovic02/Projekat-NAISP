package memtable

import (
	"fmt"

	"github.com/LukaPetkovicSV16/Projekat-NAISP/sstable"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/types"
)

type Data interface {
	Get(key string) *types.Record
	Add(record types.Record) bool
	Delete(key string) bool
	GetSortedRecordsList() []types.Record
	Clear()
	GetSize() int
}

type Memtable struct {
	MaxSize int // Max size of memtable
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

func (memtable *Memtable) Add(record types.Record) bool {
	var x = memtable.Records.Add(record)
	if memtable.MaxSize <= memtable.Records.GetSize() {
		memtable.Flush()
	}
	//println("Memtable size: ", memtable.Records.GetSize())
	return x
}

func (memtable *Memtable) Delete(key string) bool {
	return memtable.Records.Delete(key)
}

func (memtable *Memtable) Flush() {
	var records = memtable.Records.GetSortedRecordsList()
	fmt.Println(records)
	sstable.Create(records, 1)
	memtable.Records.Clear()
}
