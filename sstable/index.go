package sstable

import (
	"bytes"
	"encoding/binary"
	"os"

	"github.com/LukaPetkovicSV16/Projekat-NAISP/config"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/types"
)

type Index struct {
	KeySize uint64
	Key     string
	Offset  uint64
}

type Indexes []Index

func CreateIndexes(records []types.Record, initialOffset uint64) Indexes {
	var indexes Indexes = make([]Index, len(records))
	var offset uint64 = initialOffset
	if config.Values.Structure == "single-file" {
		for i := 0; i < len(records); i++ {
			temp := Index{
				KeySize: uint64(len(records[i].Key)),
				Key:     records[i].Key,
				Offset:  offset,
			}
			offset += uint64(len(temp.Serialize()))
		}
	}
	for i, record := range records {
		indexes[i].Key = record.Key
		indexes[i].KeySize = record.KeySize
		indexes[i].Offset = offset
		offset += uint64(len(record.Serialize()))
	}

	return indexes
}

func (index Index) Serialize() []byte {
	var serializedIndex = new(bytes.Buffer)

	binary.Write(serializedIndex, binary.LittleEndian, index.KeySize)
	binary.Write(serializedIndex, binary.LittleEndian, []byte(index.Key))
	binary.Write(serializedIndex, binary.LittleEndian, index.Offset)
	return serializedIndex.Bytes()
}

func DeserializeIndex(serializedIndex []byte) Index {
	return Index{
		KeySize: binary.LittleEndian.Uint64(serializedIndex[:8]),
		Key:     string(serializedIndex[8 : 8+binary.LittleEndian.Uint64(serializedIndex[:8])]),
		Offset:  binary.LittleEndian.Uint64(serializedIndex[8+binary.LittleEndian.Uint64(serializedIndex[:8]):]),
	}
}

// FOR MULTIPLE INDEXES
func (indexes Indexes) Serialize() []byte {
	var serializedIndexes = new(bytes.Buffer)

	for _, index := range indexes {
		binary.Write(serializedIndexes, binary.LittleEndian, index.Serialize())
	}

	return serializedIndexes.Bytes()
}

func DeserializeIndexes(serializedIndexes []byte) Indexes {
	var indexes []Index = make([]Index, 0)

	for i := 0; i < len(serializedIndexes); i += 8 + int(binary.LittleEndian.Uint64(serializedIndexes[i:i+8])) + 8 {
		indexes = append(indexes, DeserializeIndex(serializedIndexes[i:i+8+int(binary.LittleEndian.Uint64(serializedIndexes[i:i+8]))+8]))
	}

	return indexes
}

// HELPER FUNCTIONS

func readIndex(file *os.File, offset uint64, key string) *Index {
	var index Index
	file.Seek(int64(offset), 0)
	for {
		var b = make([]byte, 8)
		file.Read(b)
		index.KeySize = binary.LittleEndian.Uint64(b)
		b = make([]byte, index.KeySize)
		file.Read(b)
		index.Key = string(b)
		b = make([]byte, 8)
		file.Read(b)
		index.Offset = binary.LittleEndian.Uint64(b)
		if index.Key == key {
			return &index
		}
		if index.Key > key {
			return nil
		}
	}
}
