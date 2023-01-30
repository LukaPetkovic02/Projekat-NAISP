package main

import (
	"github.com/LukaPetkovicSV16/Projekat-NAISP/App"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/engine"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/memtable"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/skipList"
)

func main() {
	engine.CreateDataFolderStructure()
	var memtable = memtable.Init(100, skipList.NewSkipList(7))
	App.TUI(memtable)
	// var m1 = memtable.Init(100, &SkipList{})
	// var m2 = memtable.Init(100, &BTree{})
	// println(m1.Get("keyskiplist"))
	// println(m2.Get("keybtree"))
	// rand.Seed(time.Now().UnixNano())
	// level := 0
	// rand.Seed(time.Now().UnixNano())
	// for rand.Intn(2) == 1 && level < 7 {
	// 	level++
	// }
	// fmt.Println(level)

	// var rec = types.CreateRecord("key1", []byte("value1"), false)
	// var rec1 = types.CreateRecord("key2", []byte("value2"), false)
	// var rec2 = types.CreateRecord("key3", []byte("value3"), false)

	// var sl = skipList.NewSkipList(7)
	// var mem = memtable.Init(100, sl)
	// mem.Records.Add("key4", rec)
	// mem.Records.Add("key2", rec1)
	// mem.Records.Add("key1", rec2)
	// mem.Records.Add("key3", rec)
	// fmt.Println(string(mem.Get("key2").Value))
	// sl.Insert("key7", rec)
	// sl.Insert("key6", rec)
	// sl.Insert("key3", rec)
	// sl.Insert("key1", rec)
	// sl.Insert("key5", rec)
	// sl.Insert("key6", rec)
	// sl.Insert("key2", rec)
	// var sr = sl.Get("key1")
	// sl.PrintList()
	// fmt.Println(sr)
}
