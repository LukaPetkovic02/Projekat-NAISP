package types

import (
	"bytes"
	"encoding/binary"
	"hash/crc32"
	"time"
)

const (
	CRC_SIZE        = 4
	TIMESTAMP_SIZE  = 8
	TOMBSTONE_SIZE  = 1
	KEY_SIZE_SIZE   = 8
	VALUE_SIZE_SIZE = 8
)

type Record struct {
	CRC       uint32
	Timestamp uint64
	Tombstone bool
	KeySize   uint64
	ValueSize uint64
	Key       string
	Value     []byte
}

func CreateRecord(key string, value []byte, tombstone bool) Record {
	return Record{
		CRC:       crc32.ChecksumIEEE(value),
		Timestamp: uint64(time.Now().UnixMilli()),
		Tombstone: tombstone,
		KeySize:   uint64(len(key)),
		ValueSize: uint64(len(value)),
		Key:       key,
		Value:     value,
	}
}

func (record Record) Serialize() []byte {
	var serializedRecord = new(bytes.Buffer)

	binary.Write(serializedRecord, binary.LittleEndian, record.CRC)
	binary.Write(serializedRecord, binary.LittleEndian, record.Timestamp)
	binary.Write(serializedRecord, binary.LittleEndian, record.Tombstone)
	binary.Write(serializedRecord, binary.LittleEndian, record.KeySize)
	binary.Write(serializedRecord, binary.LittleEndian, record.ValueSize)
	binary.Write(serializedRecord, binary.LittleEndian, record.Key)
	binary.Write(serializedRecord, binary.LittleEndian, record.Value)
	return serializedRecord.Bytes()
}

func DeserializeRecord(serializedRecord []byte) Record {
	// TODO: Implement this function
	return Record{}
}
