package memtable

import (
	"github.com/LukaPetkovicSV16/Projekat-NAISP/types"
)

type HashMap struct {
	data map[string]*types.Record
}

func InitHash() *HashMap {
	return &HashMap{
		data: make(map[string]*types.Record),
	}
}

func (hashMap *HashMap) Get(key string) *types.Record {
	return hashMap.data[key]
}

func (hashMap *HashMap) Add(record types.Record) bool {
	hashMap.data[record.Key] = &record
	return true
}

func (hashMap *HashMap) Delete(key string) bool {
	delete(hashMap.data, key)
	return true
}

// sort by key
func (hashMap *HashMap) GetSortedRecordsList() []types.Record {
	var records []types.Record
	// sort.Slice(records, func(i, j int) bool {
	// 	return records[i].Key < records[j].Key
	// }
	return records
}

func (hashMap *HashMap) Clear() {
	hashMap.data = make(map[string]*types.Record)
}

func (hashMap *HashMap) GetSize() int {
	return len(hashMap.data)
}
