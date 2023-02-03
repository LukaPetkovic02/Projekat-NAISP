package compaction

import (
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/LukaPetkovicSV16/Projekat-NAISP/config"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/engine"
)

func LeveledCompaction(current_level int) {
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
		//ss1 := ReadAllRecords(currentLevelFiles[i-1])
		//ss2 := ReadAllRecords(currentLevelFiles[i])
		//ss3 := Merge(ss1,ss2)
		//CreateForNextLevel(ss3, current_level)
		//Delete(currentLevelFiles[i-1])
		//Delete(currentLevelFiles[i])
	}

	if current_level != config.Values.Segment_size {

	}
}
