package sstable

import (
	"bytes"
	"encoding/binary"
	"os"

	"github.com/LukaPetkovicSV16/Projekat-NAISP/types"
)

type Index struct {
	KeySize uint64
	Key     string
	Offset  uint64
}

type Indexes []Index

func CreateIndexes(records []types.Record, startOffset uint64) Indexes {
	var indexes []Index = make([]Index, len(records))
	var offset uint64 = startOffset
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

func ReadIndex(file *os.File, offset uint64) Index {
	var index Index

	file.Seek(int64(offset), 0)
	binary.Read(file, binary.LittleEndian, &index.KeySize)
	var b = make([]byte, index.KeySize)
	file.Read(b)
	index.Key = string(b)
	binary.Read(file, binary.LittleEndian, &index.Offset)

	return index
}
