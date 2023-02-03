package sstable

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/LukaPetkovicSV16/Projekat-NAISP/bloomFilter"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/config"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/engine"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/types"
)

// TODO: Get config from config file and check if single or multiple files are to be written
// TODO: Add default config in engine->constants.go
func Create(listOfRecords []types.Record) {
	if config.Values.Structure == "multiple-files" {
		writeToMultipleFiles(listOfRecords)
	} else {
		writeToSingleFile(listOfRecords)
	}
}

func CreateForLevel(listOfRecords []types.Record, level int) {
	if config.Values.Structure == "multiple-files" {
		writeToMultipleFilesForLevel(listOfRecords, level)
	} else {
		writeToSingleFileForLevel(listOfRecords, level)
	}
}

func writeToMultipleFilesForLevel(listOfRecords []types.Record, level int) {
	data := convertRecordsToBytes(listOfRecords)
	indexes := CreateIndexes(listOfRecords, 0)
	summary := CreateSummary(listOfRecords, 0)
	filter := bloomFilter.CreateBloomFilter(len(listOfRecords), config.Values.BloomFilter.Precision)
	var FILENAME = engine.GetTableNexLevelName(level) //sets same name for all files, different directories

	for _, record := range listOfRecords {
		filter.Add([]byte(record.Key))
	}
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
	fmt.Println(summary)
	file.Write(summary.Serialize())

	file, err = os.OpenFile(engine.GetBloomFilterPath(FILENAME), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	file.Write(filter.Serialize())
	file.Close()
}

func writeToSingleFileForLevel(listOfRecords []types.Record, level int) {
	var FILENAME = engine.GetTableNexLevelName(level)
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
	var summaryLength = uint64(len(summary.Serialize()))
	var indexes = CreateIndexes(listOfRecords, filterLength+summaryLength)
	var data = convertRecordsToBytes(listOfRecords)
	file.Write(data)
	file.Write(indexes.Serialize())
	file.Write(summary.Serialize())
	file.Write(filter.Serialize())
	file.Close()
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

func writeToMultipleFiles(listOfRecords []types.Record) {
	data := convertRecordsToBytes(listOfRecords)
	indexes := CreateIndexes(listOfRecords, 0)
	summary := CreateSummary(listOfRecords, 0)
	filter := bloomFilter.CreateBloomFilter(len(listOfRecords), config.Values.BloomFilter.Precision)
	var FILENAME = engine.GetTableName() //sets same name for all files, different directories

	for _, record := range listOfRecords {
		filter.Add([]byte(record.Key))
	}
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
	fmt.Println(summary)
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
	var summaryLength = uint64(len(summary.Serialize()))
	var indexes = CreateIndexes(listOfRecords, filterLength+summaryLength)
	var data = convertRecordsToBytes(listOfRecords)
	file.Write(data)
	file.Write(indexes.Serialize())
	file.Write(summary.Serialize())
	file.Write(filter.Serialize())
	file.Close()
}

func convertRecordsToBytes(listOfRecords []types.Record) []byte {
	var bytes []byte
	for _, record := range listOfRecords {
		bytes = append(bytes, record.Serialize()...)
	}
	return bytes
}

func Delete(filename string) {
	if config.Values.Structure == "multiple-files" {
		deleteMultipleFiles(filename)
	} else {
		deleteSingleFiles(filename)
	}
}

func deleteMultipleFiles(filename string) {
	e := os.Remove(engine.GetSSTablePath(filename))
	if e != nil {
		log.Fatal(e)
	}
	e = os.Remove(engine.GetIndexPath(filename))
	if e != nil {
		log.Fatal(e)
	}
	e = os.Remove(engine.GetSummaryPath(filename))
	if e != nil {
		log.Fatal(e)
	}
	e = os.Remove(engine.GetBloomFilterPath(filename))
	if e != nil {
		log.Fatal(e)
	}
}

func deleteSingleFiles(filename string) {
	e := os.Remove(engine.GetSSTablePath(filename))
	if e != nil {
		log.Fatal(e)
	}
}

// TODO: Make function for reading from sstable
// TODO: Load only part of summary file into memory

func readFromMultipleFiles(key string) *types.Record {
	var possibleFiles = CheckBloomFilter(key)
	var possibleIndexesOffsets = checkSummary(key, possibleFiles)
	var returnRecord *types.Record = nil
	var possibleIndexes = make([]Index, 0)
	for i, offset := range possibleIndexesOffsets {
		fmt.Println(i, offset)
		var file, err = os.OpenFile(engine.GetIndexPath(possibleFiles[i]), os.O_RDONLY, 0666)
		if err != nil {
			panic(err)
		}
		var index = readIndex(file, offset.Offset, key)
		possibleIndexes = append(possibleIndexes, index)
	}
	for i, index := range possibleIndexes {
		var file, err = os.OpenFile(engine.GetSSTablePath(possibleFiles[i]), os.O_RDONLY, 0666)
		if err != nil {
			panic(err)
		}
		file.Seek(int64(index.Offset), 0)
		var record = types.ReadRecord(file)
		if record.Key == key && (returnRecord == nil || record.Timestamp > returnRecord.Timestamp) {
			returnRecord = &record
		}
	}

	return returnRecord
}

func readFromSingleFile(key string) *types.Record {
	var possibleFiles = CheckBloomFilter(key)

	fmt.Println(possibleFiles)
	return nil
}

func CheckBloomFilter(key string) []string {
	files, err := os.ReadDir(engine.GetBloomDir())
	if err != nil {
		panic(err)
	}
	var possibleFiles = make([]string, 0)
	for _, file := range files {
		var filterBytes, err = ioutil.ReadFile(engine.GetBloomFilterPath(file.Name()))
		if err != nil {
			panic(err)
		}
		var filter = bloomFilter.Deserialize(filterBytes)
		var actualFilter = bloomFilter.RecreateBloomFilterBloomFilter(filter.M, filter.K, filter.Fns, filter.Podaci)
		if actualFilter.Get([]byte(key)) {
			possibleFiles = append(possibleFiles, file.Name())
		}
	}
	return possibleFiles
}
