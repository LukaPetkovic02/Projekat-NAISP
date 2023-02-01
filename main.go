package main

import (
	"github.com/LukaPetkovicSV16/Projekat-NAISP/App"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/engine"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/memtable"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/skipList"
)

func main() {
	engine.CreateDataFolderStructure()
	var sl = &skipList.SkipList{}
	sl.InitSP(10, 10, 100)
	// sl = sl.InitSP(10, 10, 100)
	var memtable = memtable.Init(100, sl)
	// // TODO: initialize LRU here and pass it to App
	App.TUI(memtable)

	// fmt.Println(engine.GetSSTableFilePath())
}
