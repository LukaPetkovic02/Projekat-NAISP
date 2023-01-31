package wal

import (
	"io"
	"log"
	"os"

	"github.com/LukaPetkovicSV16/Projekat-NAISP/engine"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/types"
	"github.com/edsrzf/mmap-go"
)

var segment_size int = 3

func Append(record types.Record) bool {
	// TODO: Open the file as memory mapped
	file, err := os.OpenFile(engine.GetCurrentWalFilePath(), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		panic(err)
	}
	current_data := ReadWalSegment(*file)

	if len(current_data) >= segment_size {
		file.Close()
		file, err = os.OpenFile(engine.GetNextWalFilePath(), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
		if err != nil {
			panic(err)
		}
		file.Close()
		return Append(record)
	}

	info, err := file.Stat()
	if err != nil {
		panic(err)
	}
	file_size := info.Size()

	err = file.Truncate(file_size + 29 + (int64)(record.KeySize) + (int64)(record.ValueSize))
	mmapFile, err := mmap.Map(file, mmap.RDWR, 0)
	if err != nil {
		log.Fatal(err)
	}

	copy(mmapFile[file_size:], record.Serialize())
	mmapFile.Unmap()
	file.Close()

	return true
}

func Clear() {
	os.RemoveAll(engine.GetWriteAheadLogDir())
	os.Mkdir(engine.GetWriteAheadLogDir(), 0755)
}

func ReadWalSegment(file os.File) []types.Record {
	log, err := io.ReadAll(&file)
	if err != nil {
		panic(err)
	}

	var ret []types.Record
	var current_position int64 = 0

	for i := 0; current_position < (int64)(len(log)) && i < segment_size; i++ {
		ret = append(ret, types.DeserializeRecord(log[current_position:]))
		current_position += int64(len(ret[len(ret)-1].Serialize()))
	}

	return ret
}
