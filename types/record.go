package types

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"time"
)

const (
	CRC_SIZE        = 4
	TIMESTAMP_SIZE  = 8
	TOMBSTONE_SIZE  = 1
	KEY_SIZE_SIZE   = 8
	VALUE_SIZE_SIZE = 8

	CRC_START        = 0
	TIMESTAMP_START  = CRC_START + CRC_SIZE
	TOMBSTONE_START  = TIMESTAMP_START + TIMESTAMP_SIZE
	KEY_SIZE_START   = TOMBSTONE_START + TOMBSTONE_SIZE
	VALUE_SIZE_START = KEY_SIZE_START + KEY_SIZE_SIZE
	KEY_START        = VALUE_SIZE_START + VALUE_SIZE_SIZE
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
	binary.Write(serializedRecord, binary.LittleEndian, []byte(record.Key))
	binary.Write(serializedRecord, binary.LittleEndian, record.Value)
	return serializedRecord.Bytes()
}

func DeserializeRecord(serializedRecord []byte) Record {
	var ret Record

	ret.CRC = binary.LittleEndian.Uint32(serializedRecord[CRC_START : CRC_START+CRC_SIZE])

	ret.KeySize = binary.LittleEndian.Uint64(serializedRecord[KEY_SIZE_START : KEY_SIZE_START+KEY_SIZE_SIZE])

	ret.ValueSize = binary.LittleEndian.Uint64(serializedRecord[VALUE_SIZE_START : VALUE_SIZE_START+VALUE_SIZE_SIZE])

	ret.Key = fmt.Sprintf("%s", serializedRecord[KEY_START:KEY_START+ret.KeySize])

	ret.Value = serializedRecord[KEY_START+ret.KeySize : KEY_START+ret.KeySize+ret.ValueSize]

	ret.Timestamp = uint64(binary.LittleEndian.Uint64(serializedRecord[TIMESTAMP_START : TIMESTAMP_START+TIMESTAMP_SIZE]))

	if serializedRecord[TOMBSTONE_START] == 1 {
		ret.Tombstone = true
	} else {
		ret.Tombstone = false
	}

	return ret
}
