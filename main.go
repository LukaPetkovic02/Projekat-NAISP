package main

import (
	//"github.com/LukaPetkovicSV16/Projekat-NAISP/App"

	"fmt"

	"github.com/LukaPetkovicSV16/Projekat-NAISP/engine"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/types"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/wal"
	//"github.com/LukaPetkovicSV16/Projekat-NAISP/memtable"
	//"github.com/LukaPetkovicSV16/Projekat-NAISP/skipList"
)

func main() {
	engine.CreateDataFolderStructure()
	//var memtable = memtable.Init(100, skipList.NewSkipList(7))
	// TODO: initialize LRU here and pass it to Appsasada
	//App.TUI(memtable)
	var x types.Record = types.CreateRecord("Joe", []byte("Bama"), false)
	wal.Append(x)
	wal.Append(x)
	wal.Append(x)
	wal.Append(x)
	wal.Append(x)
	y := wal.ReadWal()
	for i := 0; i < len(y); i++ {
		fmt.Println(y[i])
	}
}
