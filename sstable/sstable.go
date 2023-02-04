package sstable

import (
	"fmt"
	"os"

	"github.com/LukaPetkovicSV16/Projekat-NAISP/bloomFilter"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/config"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/engine"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/types"
)

func Create(listOfRecords []types.Record) {
	if config.Values.Structure == "multiple-files" {
		writeToMultipleFiles(listOfRecords)
	} else {
		writeToSingleFile(listOfRecords)
	}
}

func Read(key string) *types.Record {
	if config.Values.Structure == "multiple-files" {
		var record = readFromMultipleFiles(key)
		return record
	} else {
		var record = readFromSingleFile(key)
		return record
	}
}

func ReadAllRecordsFromTable(filename string) []types.Record {
	var result []types.Record
	if config.Values.Structure == "single-file" {
		// skip bloom filter, summary and index
	}
	// read record by record from file till EOF
	return result
}

func readFromMultipleFiles(key string) *types.Record {
	var items, err = os.ReadDir(engine.GetBloomDir())
	if err != nil {
		panic(err)
	}
	for _, item := range items {
		if item.IsDir() {
			continue
		}
		var file, err = os.Open(engine.GetBloomFilterPath(item.Name()))
		if err != nil {
			panic(err)
		}
		var filter = bloomFilter.ReadFromFile(file)
		if filter.Get([]byte(key)) {
			file, err = os.Open(engine.GetSummaryPath(item.Name()))
			if err != nil {
				panic(err)
			}
			var summary = isKeyInSummaryFile(key, file)
			if !summary {
				continue
			}
			var closestRecord = getClosestRecord(key, file)
			file, err = os.Open(engine.GetIndexPath(item.Name()))
			if err != nil {
				panic(err)
			}
			var index = readIndex(file, closestRecord.Offset, key)
			file, err = os.Open(engine.GetSSTablePath(item.Name()))
			if err != nil {
				panic(err)
			}
			file.Seek(int64(index.Offset), 0)
			var record = types.ReadRecord(file)
			return &record
		}
	}
	return nil
}
func readFromSingleFile(key string) *types.Record {
	var items, err = os.ReadDir(engine.GetTableDir())
	if err != nil {
		panic(err)
	}
	for _, item := range items {
		if item.IsDir() {
			continue
		}
		var file, err = os.Open(engine.GetSSTablePath(item.Name()))
		if err != nil {
			panic(err)
		}
		var filter = bloomFilter.ReadFromFile(file)
		if filter.Get([]byte(key)) {
			var summary = isKeyInSummaryFile(key, file)
			if !summary {
				continue
			}
			var closestRecord = getClosestRecord(key, file)
			fmt.Println("Closest: ", closestRecord)
			var index = readIndex(file, closestRecord.Offset, key)
			if index == nil {
				panic("Index is nil")
			}
			fmt.Println("Index iz ss: ", index)
			file.Seek(int64(index.Offset), 0)
			var record = types.ReadRecord(file)
			return &record
		}
	}
	return nil
}
func writeToMultipleFiles(listOfRecords []types.Record) {
	filter := bloomFilter.CreateBloomFilter(len(listOfRecords), config.Values.BloomFilter.Precision)
	summary := CreateSummary(listOfRecords, 0)
	indexes := CreateIndexes(listOfRecords, 0)
	data := types.ConvertRecordsToBytes(listOfRecords)
	for _, record := range listOfRecords {
		filter.Add([]byte(record.Key))
	}
	var FILENAME = engine.GetTableName() //sets same name for all files, different directories
	file, err := os.OpenFile(engine.GetSSTablePath(FILENAME), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	file.Write(data)

	file, err = os.OpenFile(engine.GetIndexPath(FILENAME), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	file.Write(indexes.Serialize())

	file, err = os.OpenFile(engine.GetSummaryPath(FILENAME), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	file.Write(summary.Serialize())

	file, err = os.OpenFile(engine.GetBloomFilterPath(FILENAME), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	file.Write(filter.Serialize())
	file.Close()

}
func writeToSingleFile(listOfRecords []types.Record) {
	var FILENAME = engine.GetTableName()
	var file, err = os.OpenFile(engine.GetSSTablePath(FILENAME), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	var filter = bloomFilter.CreateBloomFilter(len(listOfRecords), config.Values.BloomFilter.Precision)
	for _, record := range listOfRecords {
		filter.Add([]byte(record.Key))
	}
	var filterLength = uint64(len(filter.Serialize()))
	var summary = CreateSummary(listOfRecords, filterLength)
	fmt.Println("Summary: ", summary)
	var summaryLength = uint64(len(summary.Serialize()))
	var indexes = CreateIndexes(listOfRecords, filterLength+summaryLength)
	var data = types.ConvertRecordsToBytes(listOfRecords)
	file.Write(filter.Serialize())
	file.Write(summary.Serialize())
	file.Write(indexes.Serialize())
	file.Write(data)
	file.Close()
}
