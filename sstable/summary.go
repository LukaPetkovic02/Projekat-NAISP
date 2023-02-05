package sstable

import (
	"bytes"
	"encoding/binary"
	"os"

	"github.com/LukaPetkovicSV16/Projekat-NAISP/config"
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
	var indexes = CreateIndexes(listOfRecords, initialOffset)
	// fmt.Println("indexes: ", indexes)
	var summaryIndexes = make([]Index, 0)
	var additionalOffset uint64 = 0
	if config.Values.Structure == "single-file" {
		// fmt.Println("single-file")
		additionalOffset += 16 + uint64(len(indexes[0].Key)) + uint64(len(indexes[len(indexes)-1].Key))

		for i := 0; i < len(indexes); i++ {
			if i%config.Values.Summary.BlockSize == 0 {
				additionalOffset += uint64(len(indexes[i].Serialize()))
			}
		}
		// fmt.Println("additional offset: ", additionalOffset)
	}
	var indexOffset uint64 = initialOffset + additionalOffset
	for i := 0; i < len(indexes); i++ {
		if i%config.Values.Summary.BlockSize == 0 {
			summaryIndexes = append(summaryIndexes, Index{
				KeySize: uint64(len(indexes[i].Key)),
				Key:     indexes[i].Key,
				Offset:  indexOffset,
			})
		}
		indexOffset += uint64(len(indexes[i].Serialize()))
	}
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

func isKeyInSummaryFile(key string, file *os.File) bool {
	var summary = ReadSummaryHeader(file)
	if key >= string(summary.StartKey) && key <= string(summary.EndKey) {
		return true
	}
	return false
}

func getClosestRecord(key string, file *os.File) Index {
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

func ReadSummaryHeader(file *os.File) Summary {
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

func readFirstIndex(file *os.File) *Index {
	var index Index
	var b = make([]byte, 8)
	file.Read(b)
	index.KeySize = binary.LittleEndian.Uint64(b)
	b = make([]byte, index.KeySize)
	file.Read(b)
	index.Key = string(b)
	b = make([]byte, 8)
	file.Read(b)
	index.Offset = binary.LittleEndian.Uint64(b)
	return &index
}
