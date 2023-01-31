package engine

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

func CreateDataFolderStructure() {
	os.Mkdir(filepath.Join(DefaultDataPath, DefaultDataDir), 0755)
	os.Mkdir(filepath.Join(DefaultDataPath, DefaultDataDir, DefaultWriteAheadLogDir), 0755)
	os.Mkdir(filepath.Join(DefaultDataPath, DefaultDataDir, DefaultSSTableDir), 0755)
}

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

func GetWalFilePathBefore(current string) string {
	var start int = 0
	var end int = 0
	for i := 0; i < len(current); i++ {
		if current[i] == '_' {
			start = i + 1
		} else if current[i] == '.' {
			end = i
		}
	}
	lastFileNum, err := strconv.Atoi(current[start:end])
	if err != nil {
		panic(err)
	}
	var numberOfFiles = strconv.Itoa(lastFileNum - 1)
	if numberOfFiles == "0" {
		return ""
	}
	return filepath.Join(GetWriteAheadLogDir(), "wal_"+numberOfFiles+".log.bin")
}

func GetSSTablePath() string {
	return filepath.Join(DefaultDataPath, DefaultDataDir, DefaultSSTableDir)
}
