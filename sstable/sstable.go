package sstable

import (
	"fmt"

	"github.com/LukaPetkovicSV16/Projekat-NAISP/types"
)

// TODO: Get config from config file and check if single or multiple files are to be written
// TODO: Add default config in engine->constants.go
func Create(listOfRecords []types.Record) {
	var indexes = CreateIndex(listOfRecords)
	fmt.Println(indexes)
	// Get index bytes
	// Get index summary bytes
	// Get bloomFilter bytes
}

func ConvertRecordsToBytes(listOfRecords []types.Record) []byte {
	var bytes []byte
	for _, record := range listOfRecords {
		bytes = append(bytes, record.Serialize()...)
	}
	return bytes
}

// TODO: Make function for reading from sstable
// TODO: Load only part of summary file into memory

func Read() {
	// Check bloomFilter if key exists
	// Check index summary for offset
	// Check index for offset
	// Read from data file
}
