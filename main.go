package main

import (
	"github.com/LukaPetkovicSV16/Projekat-NAISP/App"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/engine"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/lru"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/memtable"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/skipList"
)

func main() {
	engine.CreateDataFolderStructure()
	var sl = &skipList.SkipList{}
	var LRU = lru.NewLRU(10)
	sl.InitSP(10, 10, 100)
	var memtable = memtable.Init(100, sl)
	App.TUI(memtable, LRU)

	// var time = time.Now().UnixNano()
	// fmt.Println("1_" + strconv.FormatInt(time, 10))

}
