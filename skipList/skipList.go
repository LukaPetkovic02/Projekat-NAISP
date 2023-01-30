package skipList

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/LukaPetkovicSV16/Projekat-NAISP/types"
)

type SkipList struct {
	maxHeight int
	height    int
	size      int
	head      *SkipListNode
}

type SkipListNode struct {
	Key    string
	Record types.Record
	next   []*SkipListNode
}

func NewSkipList(maxHeight int) *SkipList {

	return &SkipList{
		maxHeight: maxHeight,
		height:    0,
		size:      0,
		head:      nil,
	}

}

func NewSkipListNode(key string, record types.Record, height int) *SkipListNode {

	return &SkipListNode{
		Key:    key,
		Record: record,
		next:   make([]*SkipListNode, height+1),
	}

}

func (s *SkipList) roll() int {

	level := 0

	rand.Seed(time.Now().UnixNano())
	for rand.Intn(2) == 1 {
		level++
		if level >= s.maxHeight {
			level = s.maxHeight - 1
			break
		}
	}

	return level

}

func (s *SkipList) Get(key string) *types.Record {

	return s.search(key)

}

func (s *SkipList) search(key string) *types.Record {

	node := s.head
	if node == nil {
		return nil
	}
	for i := s.height; i >= 0; i-- {
		for node.next[i] != nil && node.next[i].Key < key {
			node = node.next[i]
		}
	}

	node = node.next[0]
	if node != nil && node.Key == key {
		return &node.Record
	}

	return nil

}

func (s *SkipList) Add(key string, record types.Record) bool {

	if key == "" || s.search(key) != nil {
		return false
	}
	level := s.roll()
	node := NewSkipListNode(key, record, s.maxHeight)

	update := make([]*SkipListNode, s.maxHeight)
	current := s.head
	if current == nil {
		s.head = node
		return true
	}
	for i := s.height; i >= 0; i-- {
		for current.next[i] != nil && current.next[i].Key < key {
			current = current.next[i]
		}
		update[i] = current
	}

	for i := 0; i <= s.height; i++ {
		node.next[i] = update[i].next[i]
		update[i].next[i] = node
	}

	s.size++
	if level > s.height {
		s.height = level
	}

	return true
}

func (s *SkipList) Update(key string, record types.Record) bool {

	node := s.head

	for i := s.height; i >= 0; i-- {
		for node.next[i] != nil && node.next[i].Key < key {
			node = node.next[i]
		}
	}

	node = node.next[0]
	if node != nil && node.Key == key {
		node.Record = record
		return true
	}

	return false

}

func (s *SkipList) Delete(key string) bool {

	update := make([]*SkipListNode, s.height+1)
	current := s.head

	for i := s.height; i >= 0; i-- {
		for current.next[i] != nil && current.next[i].Key < key {
			current = current.next[i]
		}
		update[i] = current
	}

	current = current.next[0]
	if current != nil && current.Key == key {
		current.Record.Tombstone = true
		// for i := 0; i <= s.height; i++ {
		// 	if update[i].next[i] != current {
		// 		break
		// 	}
		// 	update[i].next[i] = current.next[i]
		// }

		// s.size--
		// for s.height > 0 && s.head.next[s.height] == nil {
		// 	s.height--
		// }

		// return true

	}

	return false
}

func (s *SkipList) PrintList() {

	for i := s.height; i >= 0; i-- {
		node := s.head
		fmt.Printf("Level %d:\n", i)
		for node.next[i] != nil {
			node = node.next[i]
			fmt.Printf("\tkey: %s, value: %s\n", node.Key)
		}
	}

}

// func SkipListTest() {

// 	skiplist := NewSkipList(6)

// 	for {
// 		var key string
// 		var value string
// 		fmt.Print("Enter new key (exit to stop): ")
// 		fmt.Scan(&key)
// 		if key == "exit" {
// 			break
// 		}
// 		fmt.Print("Enter value: ")
// 		fmt.Scan(&value)
// 		inserted := skiplist.Insert(key, []byte(value))
// 		if !inserted {
// 			fmt.Println("Key '", key, "' already exists.")
// 			continue
// 		}
// 		fmt.Printf("Inserted key-value pair: (%s, %s)\n", key, value)
// 	}

// 	skiplist.PrintList()

// 	var keytoSearch string
// 	fmt.Print("Enter key to search: ")
// 	fmt.Scan(&keytoSearch)
// 	value := skiplist.Search(keytoSearch)
// 	if value != nil {
// 		fmt.Println("Value of key '", keytoSearch, "' is:", string(value))
// 	} else {
// 		fmt.Println("Key '", keytoSearch, "' is not found.")
// 	}

// 	var keyToDelete string
// 	fmt.Print("Enter key to delete: ")
// 	fmt.Scan(&keyToDelete)
// 	deleted := skiplist.Delete(keyToDelete)

// 	if deleted {
// 		fmt.Println("Deleted key '", keyToDelete, "'")
// 	} else {
// 		fmt.Println("Key '", keyToDelete, "' is not found.")
// 	}

// 	var keyToUpdate string
// 	var newValue string
// 	fmt.Print("Enter key to update: ")
// 	fmt.Scan(&keyToUpdate)
// 	fmt.Print("Enter new value: ")
// 	fmt.Scan(&newValue)
// 	updated := skiplist.Update(keyToUpdate, []byte(newValue))
// 	if updated {
// 		fmt.Println("Updated key '", keyToUpdate, "'", "with value:", newValue)
// 	} else {
// 		fmt.Println("Key '", keyToUpdate, "' is not found.")
// 	}

// 	skiplist.PrintList()

// }
