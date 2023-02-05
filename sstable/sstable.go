package sstable

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/LukaPetkovicSV16/Projekat-NAISP/bloomFilter"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/config"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/engine"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/merkleTree"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/types"
)

func Create(listOfRecords []types.Record, level int) {
	var FILENAME = engine.GetTableName(level)
	merkel := merkleTree.MerkleTree(listOfRecords)
	merkel.Serialize(FILENAME)

	if config.Values.Structure == "multiple-files" {
		writeToMultipleFiles(listOfRecords, FILENAME)
	} else {
		writeToSingleFile(listOfRecords, FILENAME)
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
			// fmt.Println("Closest: ", closestRecord)
			var index = readIndex(file, closestRecord.Offset, key)
			if index == nil {
				panic("Index is nil")
			}
			// fmt.Println("Index iz ss: ", index)
			file.Seek(int64(index.Offset), 0)
			var record = types.ReadRecord(file)
			return &record
		}
	}
	return nil
}

func writeToMultipleFiles(listOfRecords []types.Record, FILENAME string) {
	filter := bloomFilter.CreateBloomFilter(len(listOfRecords), config.Values.BloomFilter.Precision)
	summary := CreateSummary(listOfRecords, 0)
	indexes := CreateIndexes(listOfRecords, 0)
	data := types.ConvertRecordsToBytes(listOfRecords)
	for _, record := range listOfRecords {
		filter.Add([]byte(record.Key))
	}
	file, err := os.OpenFile(engine.GetSSTablePath(FILENAME), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	file.Write(data)
	file.Close()

	file, err = os.OpenFile(engine.GetIndexPath(FILENAME), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	file.Write(indexes.Serialize())
	file.Close()

	file, err = os.OpenFile(engine.GetSummaryPath(FILENAME), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	file.Write(summary.Serialize())
	file.Close()

	file, err = os.OpenFile(engine.GetBloomFilterPath(FILENAME), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	file.Write(filter.Serialize())
	file.Close()
}

func writeToSingleFile(listOfRecords []types.Record, FILENAME string) {
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
	// fmt.Println("Summary: ", summary)
	var summaryLength = uint64(len(summary.Serialize()))
	var indexes = CreateIndexes(listOfRecords, filterLength+summaryLength)
	var data = types.ConvertRecordsToBytes(listOfRecords)
	file.Write(filter.Serialize())
	file.Write(summary.Serialize())
	file.Write(indexes.Serialize())
	file.Write(data)
	file.Close()
}

func Delete(filename string) {
	e := os.Remove(engine.GetMetaDataFilePath(filename))
	if e != nil {
		log.Fatal(e)
	}

	if config.Values.Structure == "multiple-files" {
		deleteMultipleFiles(filename)
	} else {
		deleteSingleFiles(filename)
	}
}

func deleteMultipleFiles(filename string) {
	e := os.Remove(engine.GetSummaryPath(filename))
	if e != nil {
		fmt.Println("3")
		log.Fatal(e)
	}
	e = os.Remove(engine.GetBloomFilterPath(filename))
	if e != nil {
		fmt.Println("4")
		log.Fatal(e)
	}
	e = os.Remove(engine.GetSSTablePath(filename))
	if e != nil {
		fmt.Println("1")
		log.Fatal(e)
	}
	e = os.Remove(engine.GetIndexPath(filename))
	if e != nil {
		fmt.Println("2")
		log.Fatal(e)
	}
}

func deleteSingleFiles(filename string) {
	e := os.Remove(engine.GetSSTablePath(filename))
	if e != nil {
		log.Fatal(e)
	}
}

func ReadAllRecordsFromTable(filename string) []types.Record {
	file, err := os.OpenFile(engine.GetSSTablePath(filename), os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}

	var start int = 0

	if config.Values.Structure == "single-file" {
		// skip bloom filter, summary and index

		bloomFilter.ReadFromFile(file)
		ReadSummaryHeader(file)
		index := readFirstIndex(file)
		index = readIndex(file, index.Offset, index.Key)
		start = int(index.Offset)
		file.Seek(0, 0)
	}

	log, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	file.Close()
	return types.ReadRecords(log[start:])
}
