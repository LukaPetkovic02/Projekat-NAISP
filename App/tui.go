package App

import (
	"fmt"

	"github.com/LukaPetkovicSV16/Projekat-NAISP/lru"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/memtable"
)

// TODO: Add Range Scan and Get List options
func TUI(memtable *memtable.Memtable, LRU *lru.LRUCache) {
	var isRunning = true
	for isRunning {
		printMenu()
		println("Select option: ")
		var option = getUserInput()
		switch option {
		case "1":
			key, value := getKeyValue()
			HandleAdd(key, []byte(value), memtable, LRU)
		case "2":
			key := getKey()
			var record = HandleGet(key, memtable, LRU)
			fmt.Println(record)
			fmt.Println(string(record.Value))
		case "3":
			key := getKey()
			HandleDelete(key, memtable)
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
