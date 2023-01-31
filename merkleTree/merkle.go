package merkleTree

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

func Leafs(leaf []string) []*Node {

	var leafNodes []*Node

	for _, keys := range leaf {
		var hes = Hash([]byte(keys))
		leafNodes = append(leafNodes, &Node{
			Data:  hes[:],
			Left:  nil,
			Right: nil,
		})
	}

	return leafNodes
}

func Make_tree(leafNodes []*Node) *Node {

	if len(leafNodes) == 1 {
		return leafNodes[0]
	}
	if len(leafNodes)%2 == 1 {
		leafNodes = append(leafNodes, &Node{Data: []byte{}, Left: nil, Right: nil})
	}

	var parents []*Node

	for i := 0; i < len(leafNodes); i += 2 {

		var hash = append(leafNodes[i].Data, leafNodes[i+1].Data...)
		var hash2 = Hash(hash)

		parents = append(parents, &Node{Data: hash2[:], Left: leafNodes[i], Right: leafNodes[i+1]})
	}

	return Make_tree(parents)

}

func Traverse_tree(root *Node) [][]byte {

	var nodes []*Node
	var hashed_nodes [][]byte

	nodes = append(nodes, root)

	for len(nodes) > 0 {
		node := nodes[0]
		nodes = nodes[1:]

		fmt.Println(node.String())

		if node.Left != nil {
			nodes = append(nodes, node.Left)
			hashed_nodes = append(hashed_nodes, node.Left.Data)
		}
		if node.Right != nil {
			nodes = append(nodes, node.Right)
			hashed_nodes = append(hashed_nodes, node.Right.Data)
		}
	}

	return hashed_nodes

}

func Serialize_tree(hashed_nodes [][]byte) {

}