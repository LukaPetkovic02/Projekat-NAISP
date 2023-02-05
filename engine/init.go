package engine

import (
	"os"
	"path/filepath"

	"github.com/LukaPetkovicSV16/Projekat-NAISP/config"
)

func CreateDataFolderStructure() {
	os.Mkdir(filepath.Join(DefaultDataPath, DefaultDataDir), 0777)
	os.Mkdir(filepath.Join(DefaultDataPath, DefaultDataDir, DefaultWriteAheadLogDir), 0777)
	os.Mkdir(filepath.Join(DefaultDataPath, DefaultDataDir, DefaultSSTableDir), 0777)
	os.Mkdir(filepath.Join(DefaultDataPath, DefaultDataDir, DefaultMetaDataDir), 0777)
	os.Mkdir(filepath.Join(DefaultDataPath, DefaultDataDir, DEFAULT_HLL_FOLDER), 0777)
	os.Mkdir(filepath.Join(DefaultDataPath, DefaultDataDir, DEFAULT_CMS_FOLDER), 0777)
	if config.Values.Structure == "multiple-files" {
		os.Mkdir(filepath.Join(DefaultDataPath, DefaultDataDir, DEFAULT_BLOOM_FILTER_DIR), 0777)
		os.Mkdir(filepath.Join(DefaultDataPath, DefaultDataDir, DEFAULT_INDEX_FILE_DIR), 0777)
		os.Mkdir(filepath.Join(DefaultDataPath, DefaultDataDir, DEFAULT_SUMMARY_FILE_DIR), 0777)
	}
}
