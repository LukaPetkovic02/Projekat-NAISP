package App

import (
	"fmt"

	"github.com/LukaPetkovicSV16/Projekat-NAISP/lru"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/memtable"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/tokenBucket"
)

// TODO: Add Range Scan and Get List options
func TUI(memtable *memtable.Memtable, LRU *lru.LRUCache, token *tokenBucket.TokenBucket) {
	var isRunning = true
	// for _, v := range wal.ReadWal() {
	// 	HandleAdd(v.Key, v.Value, memtable, LRU)
	// }
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
			// key, value := getKeyValue()
			// HandleAdd(key, []byte(value), memtable, LRU)

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
			// var record = HandleGet(key, memtable, LRU)
			// if record == nil {
			// 	fmt.Println("Record doesn't exist")
			// 	break
			// }
			// fmt.Println(record)
			// fmt.Println(string(record.Value))
		case "3":

			if token.RequestApproval() {
				key := getKey()
				HandleDelete(key, memtable, LRU)
			}
			// key := getKey()
			// HandleDelete(key, memtable, LRU)
		case "4":
			println("Compact")
		case "5":
			isRunning = false
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
}
