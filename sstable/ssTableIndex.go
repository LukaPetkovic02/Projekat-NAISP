package sstable

import (
	"fmt"

	"github.com/LukaPetkovicSV16/Projekat-NAISP/types"
)

// TODO: write functions that return []byte that will be written to the file
func GetIndexBytes(records []types.Record) []byte {
	// For each record in records create index entry with offset and key
	var bytes []byte = nil
	var offset = 0
	for _, record := range records {
		key := record.Key
		bytes = append(bytes, key...)
		offset += len(record.Serialize())
		// bytes =
		fmt.Println(offset)
	}
	return bytes
}
