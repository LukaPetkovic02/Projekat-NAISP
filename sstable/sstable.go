package sstable

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/LukaPetkovicSV16/Projekat-NAISP/bloomFilter"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/engine"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/types"
)

// TODO: Get config from config file and check if single or multiple files are to be written
// TODO: Add default config in engine->constants.go
func Create(listOfRecords []types.Record) {
	if true {
		writeToMultipleFiles(listOfRecords)
	} else {
		writeToSingleFile(listOfRecords)
	}
}
func Read(key string) *types.Record {
	if true {
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
	filter := bloomFilter.CreateBloomFilter(len(listOfRecords), engine.DEFAULT_FILTER_PRECISION)
	file, err := os.OpenFile(engine.GetNextTableFilePath(), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	file.Write(data)

	file, err = os.OpenFile(engine.GetNextIndexFilePath(), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	file.Write(indexes.Serialize())

	file, err = os.OpenFile(engine.GetNextSummaryFilePath(), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	file.Write(summary.Serialize())

	file, err = os.OpenFile(engine.GetNextBloomFilterPath(), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	file.Write(filter.Serialize())
	file.Close()
}

func writeToSingleFile(listOfRecords []types.Record) {

}

func convertRecordsToBytes(listOfRecords []types.Record) []byte {
	var bytes []byte
	for _, record := range listOfRecords {
		bytes = append(bytes, record.Serialize()...)
	}
	return bytes
}

// TODO: Make function for reading from sstable
// TODO: Load only part of summary file into memory

func readFromMultipleFiles(key string) *types.Record {
	files, err := os.ReadDir(engine.GetBloomFilterPath())
	if err != nil {
		panic(err)
	}
	// var possibleFiles = make([]string, 0)
	for _, file := range files {
		var path = filepath.Join(engine.GetBloomFilterPath(), file.Name())
		var filterBytes, err = ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}
		var filter = bloomFilter.Deserialize(filterBytes)
		fmt.Println("Filter: ", filter.Get([]byte(key)))
		// if filter.Get([]byte(key)) {
		// 	possibleFiles = append(possibleFiles, file.Name())
		// }
		// fmt.Println("File: ", possibleFiles)
	}
	return nil
}

func readFromSingleFile(key string) *types.Record {
	return nil
}
