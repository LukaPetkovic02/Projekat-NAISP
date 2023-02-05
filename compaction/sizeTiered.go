package compaction

import (
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/LukaPetkovicSV16/Projekat-NAISP/config"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/engine"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/sstable"
)

func SizeTierCompaction(current_level int) {
	files, err := ioutil.ReadDir(engine.GetTableDir())
	if err != nil {
		log.Fatal(err)
	}

	var currentLevelFiles []string

	for _, file := range files {
		if strings.HasPrefix(file.Name(), strconv.FormatInt(int64(current_level), 10)+"_") {
			currentLevelFiles = append(currentLevelFiles, file.Name())
		}
	}

	sort.Strings(currentLevelFiles) //sortira sve fajlove
	for i := 1; i < len(currentLevelFiles); i += 2 {
		ss1 := sstable.ReadAllRecordsFromTable(currentLevelFiles[i-1])
		ss2 := sstable.ReadAllRecordsFromTable(currentLevelFiles[i])
		ss3 := Merge(ss1, ss2)

		sstable.Create(ss3, current_level+1)
		sstable.Delete(currentLevelFiles[i-1])
		sstable.Delete(currentLevelFiles[i])

	}

	if current_level+1 != int(config.Values.Lsm.MaxLevel) {
		SizeTierCompaction(current_level + 1)
	}
}

func GetWithPrefix(prefix string, page_num int, page_size int) {
	files, err := ioutil.ReadDir(engine.GetTableDir())
	if err != nil {
		log.Fatal(err)
	}
	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}
	sort.Strings(fileNames)
	for i := 0; i < len(fileNames); i++ {
		//Udje u fajl i ucita redom elemente
	}
}

// func GetAllRecords() []types.Record {
// 	files, err := ioutil.ReadDir(engine.GetTableDir())
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	list := make([]types.Record, 0)
// 	for _, file := range files {
// 		ss := sstable.ReadAllRecordsFromTable(file.Name())
// 		list = Merge(list, ss)
// 	}
// 	return list
// }
