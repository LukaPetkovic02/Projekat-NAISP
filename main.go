package main

import (
	"fmt"

	"github.com/LukaPetkovicSV16/Projekat-NAISP/App"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/engine"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/memtable"
)

func main() {
	engine.CreateDataFolderStructure()
	var memtable = memtable.Init(100, skipList.)
	// // TODO: initialize LRU here and pass it to App
	App.TUI(memtable)

	fmt.Println(engine.GetSSTableFilePath())
}
