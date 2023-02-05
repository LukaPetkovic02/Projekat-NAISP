package main

import (
	"fmt"

	"github.com/LukaPetkovicSV16/Projekat-NAISP/sstable"
)

func main() {
	// engine.CreateDataFolderStructure()
	// var LRU = lru.NewLRU(int(config.Values.Cache.Size))

	// if config.Values.Memtable.Use == "skip-list" {
	// 	var sl = &skipList.SkipList{}
	// 	sl.InitSP(int(config.Values.SkipList.MaxLevel), int(config.Values.SkipList.Height))
	// 	var memtable = memtable.Init(int(config.Values.Memtable.Size), sl)
	// 	var token = tokenBucket.Init(uint64(config.Values.TokenBucket.Size), config.Values.TokenBucket.Rate)
	// 	App.TUI(memtable, LRU, token)
	// } else {
	// 	fmt.Println("joe")
	// 	var sl = &bTree.Stablo{}
	// 	sl.InitSP(config.Values.Btree.MaxNode)
	// 	var memtable = memtable.Init(int(config.Values.Memtable.Size), sl)
	// 	var token = tokenBucket.Init(uint64(config.Values.TokenBucket.Size), config.Values.TokenBucket.Rate)
	// 	App.TUI(memtable, LRU, token)
	// }
	x := sstable.ReadAllRecordsFromTable("1_1675613797985154000.bin")
	fmt.Println(x)

	// var memtable = memtable.Init(int(config.Values.Memtable.Size), sl)
	// var token = tokenBucket.Init(uint64(config.Values.TokenBucket.Size), config.Values.TokenBucket.Rate)
	// App.TUI(memtable, LRU, token)

	// fmt.Println(cms.Deserialize(bajtovi).Encoder)
	// fmt.Println(cms.Deserialize(bajtovi).M)
	// fmt.Println(cms.Deserialize(bajtovi).Data)
	// fmt.Println(cms.Deserialize(bajtovi).Frequency([]byte("nesto")))
	// fmt.Println(cms.Deserialize(bajtovi).Frequency([]byte("nesto123")))
	// fmt.Println(cms.Deserialize(bajtovi).Frequency([]byte("nesto55")))
}
