package main

type NodeType int

const (
	NewNode NodeType = iota
	NewNodeAutoClose
	EndNode
	SpecialNode
	NoNode
)

func ReadXML(buf []byte) *Node {
	var curNode *Node

	var readedNode *Node
	var nodeType NodeType

	bufLen := len(buf)

	i := 0
	for i < bufLen {
		if buf[i] == '<' && (i == 0 || (i > 0 && buf[i-1] != '\\')) {
			i++
			readedNode, nodeType = readNodeTag(&i, buf)
			if readedNode != nil && nodeType != SpecialNode {
				if curNode != nil {
					readedNode.Parent = curNode
				}
				switch nodeType {
				case NewNode:
					curNode.Children = append(curNode.Children, readedNode)
					curNode = readedNode
					break
				case NewNodeAutoClose:
					curNode.Children = append(curNode.Children, readedNode)
					break
				case EndNode:
					curNode = curNode.Parent
					break
				}
			}
		}
		break
		i++
	}
}

func readNodeTag(i *int, buf []byte) (*Node, NodeType) {
	if buf[*i] != '<' {
		panic("Not a node beginning!")
	}

	nodeType := NoNode
	node := Node{}

	bufLen := len(buf)

	for *i < bufLen {

		*i++
	}

	return &node, nodeType
}

func readNodeContent(i *int, buf []byte) []byte {
	return nil
}
