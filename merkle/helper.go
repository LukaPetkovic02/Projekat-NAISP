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

	var cvorovi []*Node

	for _, kljucevi := range list {
		var hes = Hash([]byte(kljucevi))
		cvorovi = append(cvorovi, &Node{
			data:  hes[:],
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
	if len(listovi)%2 == 1 {
		listovi = append(listovi, &Node{data: []byte{}, left: nil, right: nil})
	}

	var roditelji []*Node

	for i := 0; i < len(listovi); i += 2 {

		var hes = append(listovi[i].data, listovi[i+1].data...)
		var hes2 = Hash(hes)

		roditelji = append(roditelji, &Node{data: hes2[:], left: listovi[i], right: listovi[i+1]})
	}

	return formiraj_stablo(roditelji)

}

func obilazak_stabla(pocetak *Node) [][]byte {

	var cvorovi []*Node
	var svi_hesevi [][]byte

	cvorovi = append(cvorovi, pocetak)

	for len(cvorovi) > 0 {
		cvor := cvorovi[0]
		cvorovi = cvorovi[1:]

		fmt.Println(cvor.String())

		if cvor.left != nil {
			cvorovi = append(cvorovi, cvor.left)
			svi_hesevi = append(svi_hesevi, cvor.left.data)
		}
		if cvor.right != nil {
			cvorovi = append(cvorovi, cvor.right)
			svi_hesevi = append(svi_hesevi, cvor.right.data)
		}
	}

	return svi_hesevi

}

func serijalizacija(hesevi [][]byte) {

}

func main() {

	var kljucevi = []string{"a", "b", "c", "d"}

	var ms = new(MerkleRoot)
	ms.root = formiraj_stablo(listovi(kljucevi))
	fmt.Println((obilazak_stabla(ms.root)))

}
