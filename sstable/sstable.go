package sstable

import (
	"fmt"

	"github.com/LukaPetkovicSV16/Projekat-NAISP/types"
)

func WriteSSTable(listOfRecords []types.Record) {
	// _, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)

	// if err != nil {
	// 	panic(err)
	// }
	for _, record := range listOfRecords {
		fmt.Println(record.Key)
	}
}
