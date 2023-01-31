package sstable

import (
	"os"

	"github.com/LukaPetkovicSV16/Projekat-NAISP/engine"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/types"
)

// TODO: Get config from config file and check if single or multiple files are to be written
// TODO: Add default config in engine->constants.go
func Create(listOfRecords []types.Record) {
	file, err := os.OpenFile(engine.GetSSTableFilePath(), os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		panic(err)
	}
	for _, record := range listOfRecords {
		file.Write(record.Serialize())
	}
}

// TODO: Make function for reading from sstable
// TODO: Load only part of summary file into memory

func Read() {

}
