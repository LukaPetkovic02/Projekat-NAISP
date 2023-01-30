package wal

import (
	"os"

	"github.com/LukaPetkovicSV16/Projekat-NAISP/engine"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/types"
)

func Append(record types.Record) bool {
	file, err := os.OpenFile(engine.GetCurrentWalFilePath(), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		panic(err)
	}
	// Write to the file
	file.Write(record.Serialize())
	file.Close()
	return true
}
