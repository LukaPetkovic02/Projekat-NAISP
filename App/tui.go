package App

import (
	//"bytes"
	"fmt"
	"io"
	"os"

	"github.com/LukaPetkovicSV16/Projekat-NAISP/cms"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/compaction"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/engine"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/hll"

	"github.com/LukaPetkovicSV16/Projekat-NAISP/lru"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/memtable"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/tokenBucket"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/wal"
)

func meniCMS() {
	for true {
		println("Enter option:")
		var cmsoption = getUserInput()
		var Cms cms.CountMinSketch
		Cms = *cms.CreateCMS(0.01, 0.01)
		switch cmsoption {
		case "1":
			println("Enter cms file name:")
			key := getUserInput()
			cmsfilename := engine.GetCMSName(key)
			file, err := os.OpenFile(engine.GetCMSPath(cmsfilename), os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				panic(err)
			}
			Cms = *cms.CreateCMS(0.01, 0.01)
			file.Write(Cms.Serialize())
			file.Close()

		case "2":
			println("Enter cms file name:")
			key := getUserInput()
			cmsfilename := engine.GetCMSName(key)
			file, err := os.OpenFile(engine.GetCMSPath(cmsfilename), os.O_RDONLY|os.O_CREATE, 0666)
			if err != nil {
				panic(err)
			}
			log, err := io.ReadAll(file)
			if err != nil {
				panic(err)
			}
			//Cms:=cms.CreateCMS(0.01,0.01)
			Cms = *cms.Deserialize(log)
			file.Close()

		case "3":
			println("Enter new element:")
			key := getUserInput()
			Cms.Add([]byte(key))

		case "4":
			println("Enter key:")
			key := getUserInput()
			fmt.Println("Frequency in cms:", Cms.Frequency([]byte(key)))

		case "5":
			println("Enter cms file name:")
			key := getUserInput()
			cmsfilename := engine.GetCMSName(key)
			file, err := os.OpenFile(engine.GetCMSPath(cmsfilename), os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				panic(err)
			}
			file.Write(Cms.Serialize())
			file.Close()

		case "6":
			println("Enter cms file name:")
			key := getUserInput()
			cmsfilename := engine.GetCMSName(key)
			e := os.Remove(engine.GetCMSPath(cmsfilename))
			if e != nil {
				fmt.Println("fajl ne postoji")
			} else {
				file, err := os.OpenFile(engine.GetCMSPath(cmsfilename), os.O_WRONLY|os.O_CREATE, 0666)
				if err != nil {
					panic(err)
				}
				file.Close()
			}

		case "7":
			return
		default:
			println("Wrong input!")
		}
	}
}

func HLL_Menu() {

	var hll_option = getUserInput()

	HLL := *hll.NewHLL(5)
	//main_key := "a"

	switch hll_option {
	// Create new HLL
	case "1":
		fmt.Println("Enter HLL key: ")
		key := getUserInput()
		file, err := os.OpenFile(engine.GetHLLPath(key+".bin"), os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			panic(err)
		}

		new_hll := hll.NewHLL(5)
		bytes := new_hll.Serialize()
		file.Write(bytes)
		file.Close()
	// get hll from file
	case "2":
		println("Enter hll file name:")
		key := getUserInput()
		hllfilename := engine.GetHLLName(key)
		file, err := os.OpenFile(engine.GetHLLPath(hllfilename), os.O_RDONLY|os.O_CREATE, 0666)
		if err != nil {
			panic(err)
		}
		log, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}
		HLL = *hll.DeSerialize(log)
		file.Close()
		// key := getUserInput()
		// file, err := os.OpenFile(engine.GetHLLPath(key+".bin"), os.O_RDONLY, 0666)
		// if err != nil {
		// 	panic(err)
		// }
		// hl,err := io.ReadAll(file)
		// HLL = hll.DeSerialize(hl)
		// file.Close()

	// Add element to HLL
	case "3":
		fmt.Println("Enter new element: ")
		element := getUserInput()
		HLL.Add([]byte(element))

	// Get estimate
	case "4":
		fmt.Println("Estimate is: ", HLL.Estimate())

	// Write HLL to file
	case "5":
		println("Enter hll file name:")
		key := getUserInput()
		hllfilename := engine.GetHLLName(key)
		file, err := os.OpenFile(engine.GetHLLPath(hllfilename), os.O_WRONLY, 0666)
		if err != nil {
			panic(err)
		}
		file.Write(HLL.Serialize())
		file.Close()
	// Delete HLL from file
	case "6":
		println("Enter cms file name:")
		key := getUserInput()
		hllfilename := engine.GetHLLName(key)
		e := os.Remove(engine.GetHLLPath(hllfilename))
		if e != nil {
			fmt.Println("fajl ne postoji")
		} else {
			file, err := os.OpenFile(engine.GetHLLPath(hllfilename), os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				panic(err)
			}
			file.Close()
		}
	case "7":
		return
	default:
		fmt.Println("Wrong choice")

	}
}

// TODO: Add Range Scan and Get List options
func TUI(memtable *memtable.Memtable, LRU *lru.LRUCache, token *tokenBucket.TokenBucket) {
	var isRunning = true
	for _, v := range wal.ReadWal() {
		HandleAdd(v.Key, v.Value, memtable, LRU)
	}
	for isRunning {
		printMenu()
		println("Select option: ")
		var option = getUserInput()
		switch option {
		case "1":

			if token.RequestApproval() {
				key, value := getKeyValue()
				HandleAdd(key, []byte(value), memtable, LRU)
			}

		case "2":
			if token.RequestApproval() {
				key := getKey()
				var record = HandleGet(key, memtable, LRU)
				if record == nil {
					fmt.Println("Record doesn't exist")
					break
				}
				fmt.Println(record)
				fmt.Println(string(record.Value))
			}

		case "3":

			if token.RequestApproval() {
				key := getKey()
				HandleDelete(key, memtable, LRU)
			}

		case "4":
			println("Compact")
			compaction.SizeTierCompaction(1)
		case "5":
			isRunning = false
		case "6":
			meniCMS()
		// 	var hll_option = getUserInput()

		// 	HLL := hll.NewHLL(5)
		// 	main_key := "a"

		// 	switch hll_option {
		// 		case "1":
		// 			fmt.Println("Enter HLL key: ")
		// 			key := getUserInput()
		// 			file, err = os.OpenFile(engine.GetHLLPath(key + ".bin"), os.O_WRONLY|os.O_CREATE, 0666)
		// 			if err != nil {
		// 				panic(err)
		// 			}

		// 			new_hll := hll.NewHLL(5)
		// 			bytes = new_hll.Serialize()
		// 			file.write(bytes)
		// 			file.Close()
		// 		// get hll from file
		// 		case "2":
		// 			key := getUserInput()
		// 			file, err = os.OpenFile(engine.GetHLLPath(key + ".bin"), os.O_RDONLY|os.O_CREATE, 0666)
		// 			if err != nil {
		// 				panic(err)
		// 			}

		// 			file.Close()

		// 		// Add element to HLL
		// 		case "3":
		// 			fmt.Println("Enter new element: ")
		// 			element := getUserInput()
		// 			HLL.Add([]byte(element))

		// 		// Get estimate
		// 		case "4":
		// 			fmt.Println("Estimate is: ",HLL.Estimate())

		// 		// Write HLL to file
		// 		case "5":
		// 			println("Enter cms file name:")
		// 			key:=getUserInput()
		// 			cmsfilename:=engine.GetCMSName(key)
		// 			file, err = os.OpenFile(engine.GetCMSPath(cmsfilename), os.O_WRONLY|os.O_CREATE, 0666)
		// 			if err != nil {
		// 				panic(err)
		// 			}
		// 			file.Write(Cms.Serialize())
		// 			file.Close()
		// 		default:
		// 			fmt.Println("asdads")

		// 	}

		default:
			println("Invalid option")
		}

	}

}

func getKey() string {
	fmt.Println("Enter key: ")
	return getUserInput()
}

func getKeyValue() (string, string) {
	fmt.Println("Enter key: ")
	var key = getUserInput()
	fmt.Println("Enter value: ")
	var value = getUserInput()
	return key, value
}

func getUserInput() string {
	var input string
	fmt.Scanln(&input)
	return input
}

// TOOD: Add Range Scan and Get List options
func printMenu() {
	println("1. Add record")
	println("2. Get record")
	println("3. Delete record")
	println("4. Compact")
	println("5. Exit")
	println("6. CMS")
	println("7. HyperLogLog")
}

func printHLLMenu() {
	println("1. Create new HLL")
	println("2. Get HLL from file")
	println("3. Add element to HLL")
	println("4. Get estimate")
	println("5. Write HLL to file")
	println("6. Delete HLL from file")
	println("7. Exit to main menu")
}

func printCMSMenu() {
	println("1. Create new CMS")
	println("2. Get CMS from file")
	println("3. Add element to CMS")
	println("4. Check frequency of element")
	println("5. Write CMS to selected file")
	println("6. Delete CMS from selected file")
	println("7. Exit to main menu")
}
