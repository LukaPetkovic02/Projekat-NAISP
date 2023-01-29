package merkel

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

type MerkleRoot struct {
	Root *Node
}

func (mr *MerkleRoot) String() string {
	return mr.Root.String()
}

type Node struct {
	Data  []byte
	Left  *Node
	Right *Node
}

func (n *Node) String() string {
	return hex.EncodeToString(n.Data[:])
}

func Hash(Data []byte) [20]byte {
	return sha1.Sum(Data)
}

func Listovi(list []string) []*Node {

	var cvorovi []*Node

	for _, kljucevi := range list {
		var hes = Hash([]byte(kljucevi))
		cvorovi = append(cvorovi, &Node{
			Data:  hes[:],
			Left:  nil,
			Right: nil,
		})
	}

	return cvorovi
}

func Formiraj_stablo(listovi []*Node) *Node {

	if len(listovi) == 1 {
		return listovi[0]
	}
	if len(listovi)%2 == 1 {
		listovi = append(listovi, &Node{Data: []byte{}, Left: nil, Right: nil})
	}

	var roditelji []*Node

	for i := 0; i < len(listovi); i += 2 {

		var hes = append(listovi[i].Data, listovi[i+1].Data...)
		var hes2 = Hash(hes)

		roditelji = append(roditelji, &Node{Data: hes2[:], Left: listovi[i], Right: listovi[i+1]})
	}

	return Formiraj_stablo(roditelji)

}

func Obilazak_stabla(pocetak *Node) [][]byte {

	var cvorovi []*Node
	var svi_hesevi [][]byte

	cvorovi = append(cvorovi, pocetak)

	for len(cvorovi) > 0 {
		cvor := cvorovi[0]
		cvorovi = cvorovi[1:]

		fmt.Println(cvor.String())

		if cvor.Left != nil {
			cvorovi = append(cvorovi, cvor.Left)
			svi_hesevi = append(svi_hesevi, cvor.Left.Data)
		}
		if cvor.Right != nil {
			cvorovi = append(cvorovi, cvor.Right)
			svi_hesevi = append(svi_hesevi, cvor.Right.Data)
		}
	}

	return svi_hesevi

}

func Serijalizacija(hesevi [][]byte) {

}
