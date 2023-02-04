package wal

import (
	"io"
	"os"
	"path/filepath"

	"github.com/LukaPetkovicSV16/Projekat-NAISP/config"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/engine"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/types"
	"github.com/edsrzf/mmap-go"
)

func Append(record types.Record) bool {
	// TODO: Open the file as memory mapped
	file, err := os.OpenFile(engine.GetCurrentWalFilePath(), os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}
	current_data := ReadWalSegment(*file)
	if len(current_data) == config.Values.WalSegment {
		file.Close()
		file, err = os.OpenFile(engine.GetNextWalFilePath(), os.O_RDWR|os.O_CREATE, 0777)
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
		panic(err)
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

	for i := 0; current_position < (int64)(len(log)) && i < config.Values.WalSegment; i++ {
		ret = append(ret, types.DeserializeRecord(log[current_position:]))
		current_position += int64(len(ret[len(ret)-1].Serialize()))
	}

	return ret
}

func ReadWal() []types.Record {
	var current_filename string = filepath.Join(engine.GetWriteAheadLogDir(), "wal_1.log.bin")
	var ret []types.Record
	for current_filename != "" {
		file, err := os.OpenFile(current_filename, os.O_RDWR, 0777)
		if err != nil {
			break
		}
		current_data := ReadWalSegment(*file)
		for i := 0; i < len(current_data); i++ {
			ret = append(ret, current_data[i])
		}
		file.Close()

		current_filename = engine.GetNextWalFile(current_filename)

	}
	return ret
}
