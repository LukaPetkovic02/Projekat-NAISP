package main

import (
	"fmt"

	"github.com/LukaPetkovicSV16/Projekat-NAISP/engine"
)

func main() {
	engine.CreateDataFolderStructure()
	// var memtable = memtable.Init(100, skipList.NewSkipList(7))
	// // TODO: initialize LRU here and pass it to App
	// App.TUI(memtable)

	fmt.Println(engine.GetSSTableFilePath())
}
