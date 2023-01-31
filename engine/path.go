package engine

func GetSSTableFilePath() string {
	// var directoryPath = GetSSTableDirPath()
	// var files, _ = ioutil.ReadDir(directoryPath)
	// var level = 1 // All files are added to same level i think it is level 1?
	// var maxIndex = "0"
	// // fmt.Println(files)
	// for _, file := range files {
	// 	// fmt.Println(file.Name())
	// 	var temp = strings.Split(file.Name(), "_")
	// 	index := temp[1]
	// 	if index > maxIndex {
	// 		maxIndex = index
	// 	}

	// }
	// var nextIndex, _ = strconv.Atoi(maxIndex)
	// var filename = fmt.Sprintf("%d_%d_sstable.bin", level, nextIndex)
	// return path.Join(directoryPath, filename)
}
