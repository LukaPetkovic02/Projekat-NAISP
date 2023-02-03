package engine

import (
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func GetWriteAheadLogDir() string {
	return filepath.Join(DefaultDataPath, DefaultDataDir, DefaultWriteAheadLogDir)
}

func GetCurrentWalFilePath() string {
	files, err := ioutil.ReadDir(GetWriteAheadLogDir())
	if err != nil {
		panic(err)
	}
	var numberOfFiles = strconv.Itoa(len(files))
	if numberOfFiles == "0" {
		numberOfFiles = "1"
	}
	return filepath.Join(GetWriteAheadLogDir(), "wal_"+numberOfFiles+".log.bin")
}

func GetNextWalFilePath() string {
	files, err := ioutil.ReadDir(GetWriteAheadLogDir())
	if err != nil {
		panic(err)
	}
	var numberOfFiles = strconv.Itoa(len(files) + 1)

	return filepath.Join(GetWriteAheadLogDir(), "wal_"+numberOfFiles+".log.bin")
}

func GetNextWalFile(current string) string {
	var start int = 0
	var end int = 0
	for i := 0; i < len(current); i++ {
		if current[i] == '_' {
			start = i + 1
		} else if current[i] == '.' && start != 0 {
			end = i
			break
		}
	}
	lastFileNum, err := strconv.Atoi(current[start:end])
	if err != nil {
		panic(err)
	}
	var numberOfFiles = strconv.Itoa(lastFileNum + 1)

	return filepath.Join(GetWriteAheadLogDir(), "wal_"+numberOfFiles+".log.bin")
}

func GetDataDir() string {
	return filepath.Join(DefaultDataPath, DefaultDataDir)
}

func GetMetaDataFilePath() string {
	_, err := ioutil.ReadDir(GetDataDir())
	if err != nil {
		panic(err)
	}

	return filepath.Join(GetDataDir(), "meta", "Metadata.txt")
}

// SSTable
func GetTableName() string {
	return "1_" + strconv.FormatInt(time.Now().UnixNano(), 10) + ".bin"
}

// za trenutni nivo racuna ime fajla u narednom nivou
func GetTableNexLevelName(current_level int) string {
	return strconv.FormatInt(int64(current_level+1), 10) + "_" + strconv.FormatInt(time.Now().UnixNano(), 10) + ".bin"
}

// za trenutno ime fajla racuna trenutni nivo ne znam da li ce mi trebati
func GetCurrentLevel(filename string) int {
	split := strings.Split(filename, ",")
	ret, _ := strconv.Atoi(split[0])
	return ret
}

func GetSSTablePath(filename string) string {
	return filepath.Join(GetDataDir(), DefaultSSTableDir, filename)
}

func GetIndexPath(filename string) string {
	return filepath.Join(GetDataDir(), DEFAULT_INDEX_FILE_DIR, filename)
}

func GetBloomFilterPath(filename string) string {
	return filepath.Join(GetDataDir(), DEFAULT_BLOOM_FILTER_DIR, filename)
}

func GetSummaryPath(filename string) string {
	return filepath.Join(GetDataDir(), DEFAULT_SUMMARY_FILE_DIR, filename)
}

// DIRS

func GetBloomDir() string {
	return filepath.Join(DefaultDataPath, DefaultDataDir, DEFAULT_BLOOM_FILTER_DIR)
}
