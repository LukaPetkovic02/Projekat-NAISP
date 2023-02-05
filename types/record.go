package types

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"os"
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

func ReadRecord(file *os.File) Record {
	var b = make([]byte, CRC_SIZE)
	file.Read(b)
	var record Record
	record.CRC = binary.LittleEndian.Uint32(b)
	b = make([]byte, TIMESTAMP_SIZE)
	file.Read(b)
	record.Timestamp = binary.LittleEndian.Uint64(b)
	b = make([]byte, TOMBSTONE_SIZE)
	file.Read(b)
	if b[0] == 1 {
		record.Tombstone = true
	} else {
		record.Tombstone = false
	}
	b = make([]byte, KEY_SIZE_SIZE)
	file.Read(b)
	record.KeySize = binary.LittleEndian.Uint64(b)
	b = make([]byte, VALUE_SIZE_SIZE)
	file.Read(b)
	record.ValueSize = binary.LittleEndian.Uint64(b)
	b = make([]byte, int(record.KeySize))
	file.Read(b)
	record.Key = string(b)
	b = make([]byte, int(record.ValueSize))
	file.Read(b)
	record.Value = b
	return record
}

func ConvertRecordsToBytes(listOfRecords []Record) []byte {
	var bytes []byte
	for _, record := range listOfRecords {
		bytes = append(bytes, record.Serialize()...)
	}
	return bytes
}

// Ucitava niz rekorda iz niza bajtova, pogdno za sstable i druge strukture
func ReadRecords(data []byte) []Record {
	var ret []Record
	var current_position = 0

	for i := 0; current_position < len(data); i++ {
		ret = append(ret, DeserializeRecord(data[current_position:]))
		current_position += len(ret[len(ret)-1].Serialize())
	}

	return ret
}
