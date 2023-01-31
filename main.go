package main

import (
	//"github.com/LukaPetkovicSV16/Projekat-NAISP/App"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/engine"
	//"github.com/LukaPetkovicSV16/Projekat-NAISP/memtable"
	//"github.com/LukaPetkovicSV16/Projekat-NAISP/skipList"
)

func main() {
	engine.CreateDataFolderStructure()
	//var memtable = memtable.Init(100, skipList.NewSkipList(7))
	// TODO: initialize LRU here and pass it to Appsasada
	//App.TUI(memtable)
}
