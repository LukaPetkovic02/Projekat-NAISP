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

func GetSSTablePath() string {
	return filepath.Join(DefaultDataPath, DefaultDataDir, DefaultSSTableDir)
}
