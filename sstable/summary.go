package sstable

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"

	"github.com/LukaPetkovicSV16/Projekat-NAISP/engine"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/types"
)

// TODO: write functions that return []byte that will be written to the file

type Summary struct {
	StartKeySize uint64
	EndKeySize   uint64
	StartKey     string
	EndKey       string
	Indexes      Indexes
}

func CreateSummary(listOfRecords []types.Record, initialOffset uint64) Summary {
	//TODO: Fix indexes to only have part of the keys
	var indexes = CreateIndexes(listOfRecords, initialOffset)
	fmt.Println(indexes)
	var summaryIndexes = make([]Index, 0)
	var indexOffset uint64 = 0
	for i := 0; i < len(indexes); i++ {
		if i%2 == 0 { // TODO: change 2 to something from config file (2 is block size of index summary)
			var temp = Index{
				KeySize: indexes[i].KeySize,
				Key:     indexes[i].Key,
				Offset:  indexOffset,
			}
			summaryIndexes = append(summaryIndexes, temp)
		}
		indexOffset += uint64(len(indexes[i].Serialize()))
	}
	fmt.Println(summaryIndexes)
	return Summary{
		StartKeySize: uint64(len(indexes[0].Key)),
		EndKeySize:   uint64(len(indexes[len(indexes)-1].Key)),
		StartKey:     indexes[0].Key,
		EndKey:       indexes[len(indexes)-1].Key,
		Indexes:      summaryIndexes,
	}

}

func (summary Summary) Serialize() []byte {
	var serializedSummary = new(bytes.Buffer)

	binary.Write(serializedSummary, binary.LittleEndian, uint64(len(summary.StartKey)))
	binary.Write(serializedSummary, binary.LittleEndian, uint64(len(summary.EndKey)))
	binary.Write(serializedSummary, binary.LittleEndian, []byte(summary.StartKey))
	binary.Write(serializedSummary, binary.LittleEndian, []byte(summary.EndKey))
	binary.Write(serializedSummary, binary.LittleEndian, summary.Indexes.Serialize())

	return serializedSummary.Bytes()
}

func DeserializeSummary(serializedSummary []byte) Summary {
	var startKeySize = binary.LittleEndian.Uint64(serializedSummary[:8])
	var endKeySize = binary.LittleEndian.Uint64(serializedSummary[8:16])
	var startKey = string(serializedSummary[16 : 16+startKeySize])
	var endKey = string(serializedSummary[16+startKeySize : 16+startKeySize+endKeySize])

	return Summary{
		StartKeySize: startKeySize,
		EndKeySize:   endKeySize,
		StartKey:     startKey,
		EndKey:       endKey,
		Indexes:      DeserializeIndexes(serializedSummary[16+startKeySize+endKeySize:]),
	}
}

func isKeyInSummaryFile(key string, filename string) bool {
	file, err := os.OpenFile(engine.GetSummaryPath(filename), os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}
	var summary = ReadSummaryHeader(file)
	if key >= string(summary.StartKey) && key <= string(summary.EndKey) {
		return true
	}
	return false

}

func getClosestRecord(key string, filename string) Index {
	file, err := os.OpenFile(engine.GetSummaryPath(filename), os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}
	file.Seek(0, 0)
	var _ = ReadSummaryHeader(file) // Used just to skip the header
	var returnIndex Index
	for {
		var tempIndex Index
		var b = make([]byte, 8)
		file.Read(b)
		tempIndex.KeySize = binary.LittleEndian.Uint64(b)
		b = make([]byte, tempIndex.KeySize)
		file.Read(b)
		tempIndex.Key = string(b)
		b = make([]byte, 8)
		file.Read(b)
		tempIndex.Offset = binary.LittleEndian.Uint64(b)
		if tempIndex.Key == key {
			returnIndex = tempIndex
			break
		}
		if tempIndex.Key > key {
			break
		}
		returnIndex = tempIndex
	}
	return returnIndex
}

func checkSummary(key string, possibleFileNames []string) []Index {
	// var files = make([]string, 0)
	var possibleIndexes = make([]Index, 0)
	for _, filename := range possibleFileNames {
		if isKeyInSummaryFile(key, filename) {
			// files = append(files, filename)
			var index = getClosestRecord(key, filename)
			possibleIndexes = append(possibleIndexes, index)
		}
	}
	return possibleIndexes
}

func ReadSummaryHeader(file *os.File) Summary {
	file.Seek(0, 0)
	var b = make([]byte, 8)
	file.Read(b)
	var startKeySize = binary.LittleEndian.Uint64(b)
	b = make([]byte, 8)
	file.Read(b)
	var endKeySize = binary.LittleEndian.Uint64(b)
	b = make([]byte, startKeySize)
	file.Read(b)
	var startKey = string(b)
	b = make([]byte, endKeySize)
	file.Read(b)
	var endKey = string(b)
	return Summary{
		StartKeySize: startKeySize,
		EndKeySize:   endKeySize,
		StartKey:     startKey,
		EndKey:       endKey,
	}
}
