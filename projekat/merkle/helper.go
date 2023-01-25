package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

type MerkleRoot struct {
	root *Node
}

func (mr *MerkleRoot) String() string {
	return mr.root.String()
}

type Node struct {
	data  []byte
	left  *Node
	right *Node
}

func (n *Node) String() string {
	return hex.EncodeToString(n.data[:])
}

func Hash(data []byte) [20]byte {
	return sha1.Sum(data)
}

func listovi(list []string) []*Node {

	var hesirani [][20]byte
	var cvorovi []*Node

	for _, kljucevi := range list {
		hesirani = append(hesirani, Hash([]byte(kljucevi)))
	}

	for _, hes := range hesirani {

		cvorovi = append(cvorovi, &Node{
			data:  hes[:],
			left:  nil,
			right: nil,
		})
	}

	if len(cvorovi)%2 == 1 {
		cvorovi = append(cvorovi, &Node{
			data:  []byte{},
			left:  nil,
			right: nil,
		})
	}

	return cvorovi
}

func formiraj_stablo(listovi []*Node) *Node {

	if len(listovi) == 1 {
		return listovi[0]
	}

	var roditelji []*Node

	for i := 0; i < len(listovi); i += 2 {

		var hes = append(listovi[i].data, listovi[i+1].data...)
		var hes2 = Hash(hes)

		roditelji = append(roditelji, &Node{
			data:  hes2[:],
			left:  listovi[i],
			right: listovi[i+1],
		})
	}

	return formiraj_stablo(roditelji)

}

func obilazak(koren *Node) {

	if koren.left == nil || koren.right == nil {
		return
	} else {
		fmt.Println(koren.left.String())
		fmt.Println(koren.right.String())
	}

	obilazak(koren.left)
	obilazak(koren.right)

}

func main() {

	var kljucevi = []string{"abc", "asd", "bcd", "asc"}

	var ms = new(MerkleRoot)
	ms.root = formiraj_stablo(listovi(kljucevi))
	obilazak(ms.root)

}
