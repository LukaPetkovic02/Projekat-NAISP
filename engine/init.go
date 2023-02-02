package engine

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

func CreateDataFolderStructure() {
	os.Mkdir(filepath.Join(DefaultDataPath, DefaultDataDir), 0777)
	os.Mkdir(filepath.Join(DefaultDataPath, DefaultDataDir, DefaultWriteAheadLogDir), 0777)
	os.Mkdir(filepath.Join(DefaultDataPath, DefaultDataDir, DefaultSSTableDir), 0777)
	os.Mkdir(filepath.Join(DefaultDataPath, DefaultDataDir, DefaultMetaDataDir), 0777)
	os.Mkdir(filepath.Join(DefaultDataPath, DefaultDataDir, DEFAULT_BLOOM_FILTER_DIR), 0777)
	os.Mkdir(filepath.Join(DefaultDataPath, DefaultDataDir, DEFAULT_INDEX_FILE_DIR), 0777)
	os.Mkdir(filepath.Join(DefaultDataPath, DefaultDataDir, DEFAULT_SUMMARY_FILE_DIR), 0777)
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

func GetSSTablePath() string {
	return filepath.Join(DefaultDataPath, DefaultDataDir, DefaultSSTableDir)
}

func GetNextFileIndex(dir string) string {
	files, err := ioutil.ReadDir(GetSSTablePath())
	if err != nil {
		panic(err)
	}
	if len(files) == 0 {
		return "1"
	}
	return strconv.Itoa(len(files))
}

func GetNextTableFilePath() string {
	return filepath.Join(GetSSTablePath(), "1_"+GetNextFileIndex(GetSSTablePath())+"_table.bin")
}

func GetBloomFilterPath() string {
	return filepath.Join(DefaultDataPath, DefaultDataDir, DEFAULT_BLOOM_FILTER_DIR)
}

func GetSummaryFilePath() string {
	return filepath.Join(DefaultDataPath, DefaultDataDir, DEFAULT_SUMMARY_FILE_DIR)
}

func GetIndexFilePath() string {
	return filepath.Join(DefaultDataPath, DefaultDataDir, DEFAULT_INDEX_FILE_DIR)
}

func GetDataPath() string {
	return filepath.Join(DefaultDataPath, DefaultDataDir, DefaultMetaDataDir)
}

func GetMetaDataFilePath() string {
	_, err := ioutil.ReadDir(GetDataPath())
	if err != nil {
		panic(err)
	}

	return filepath.Join(GetDataPath(), "Metadata.txt")
}

func GetNextIndexFilePath() string {
	return filepath.Join(GetIndexFilePath(), "1_"+GetNextFileIndex(GetIndexFilePath())+"_table.bin")
}

func GetNextBloomFilterPath() string {
	return filepath.Join(GetBloomFilterPath(), "1_"+GetNextFileIndex(GetBloomFilterPath())+"_table.bin")
}

func GetNextSummaryFilePath() string {
	return filepath.Join(GetSummaryFilePath(), "1_"+GetNextFileIndex(GetSummaryFilePath())+"_table.bin")
}

func GetNextSummaryFilePathByFilename(filename string) string {
	return filepath.Join(GetSummaryFilePath(), filename)
}
