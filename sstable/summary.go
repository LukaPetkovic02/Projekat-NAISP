package sstable

import (
	"bytes"
	"encoding/binary"
	"os"

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

	return Summary{
		StartKeySize: uint64(len(indexes[0].Key)),
		EndKeySize:   uint64(len(indexes[len(indexes)-1].Key)),
		StartKey:     indexes[0].Key,
		EndKey:       indexes[len(indexes)-1].Key,
		Indexes:      indexes,
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
	var startKeySize = make([]byte, 8)
	var endKeySize = make([]byte, 8)
	var startKey []byte
	var endKey []byte
	file.Seek(0, 0)
	file.Read(startKeySize)
	file.Read(endKeySize)

	startKey = make([]byte, binary.LittleEndian.Uint64(startKeySize))
	endKey = make([]byte, binary.LittleEndian.Uint64(endKeySize))

	file.Read(startKey)
	file.Read(endKey)

	if key >= string(startKey) && key <= string(endKey) {
		return true
	}
	return false

}
