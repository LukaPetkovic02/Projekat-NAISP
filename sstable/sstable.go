package sstable

import (
	"fmt"
	"os"

	"github.com/LukaPetkovicSV16/Projekat-NAISP/bloomFilter"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/engine"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/types"
)

// TODO: Get config from config file and check if single or multiple files are to be written
// TODO: Add default config in engine->constants.go
func Create(listOfRecords []types.Record) {
	var indexes = CreateIndexes(listOfRecords)
	fmt.Println(indexes)
	fmt.Println(DeserializeIndexes(indexes.Serialize()))
	var summary = CreateSummary(indexes)
	fmt.Println(summary)
	fmt.Println(summary.Serialize())
	var bloomFilter = bloomFilter.CreateBloomFilter(len(listOfRecords), 0.01)
	for _, record := range listOfRecords {
		bloomFilter.Add([]byte(record.Key))
	}
	// fmt.Println(indexes.Serialize())
	// fmt.Println(DeserializeIndexes(indexes.Serialize()))
	WriteToMultipleFiles(bloomFilter.Serialize(), summary.Serialize(), indexes.Serialize(), ConvertRecordsToBytes(listOfRecords))
	// Get index bytes
	// Get index summary bytes
	// Get bloomFilter bytes
}

func WriteToMultipleFiles(filter []byte, summary []byte, indexes []byte, data []byte) {
	file, err := os.OpenFile(engine.GetNextTableFilePath(), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	file.Write(data)

	file, err = os.OpenFile(engine.GetNextIndexFilePath(), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	file.Write(indexes)

	file, err = os.OpenFile(engine.GetNextSummaryFilePath(), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	file.Write(summary)
	isKeyInSummaryFile("a", file)

	file, err = os.OpenFile(engine.GetNextBloomFilterPath(), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	file.Write(filter)
	file.Close()
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
