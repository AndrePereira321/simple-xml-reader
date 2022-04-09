package main

import (
	"bytes"
)

type NodeHandler func(node *Node)

type Attr struct {
	Name  string
	Value string
}

type Node struct {
	Name     string
	Attrs    []*Attr
	Content  []byte
	Children []*Node
	Parent   *Node
}

func (this *Node) ForEachNode(handler NodeHandler) {
	handler(this)
	for _, n := range this.Children {
		n.ForEachNode(handler)
	}
}

func (node *Node) ForEachNodeReverse(handler NodeHandler) {
	for _, n := range node.Children {
		n.ForEachNodeReverse(handler)
	}
	handler(node)
}

func (n *Node) ToString() string {
	return n.toByteBuffer("").String()
}

func (n *Node) ToXMLByteArray() []byte {
	return n.toByteBuffer("<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?>\n").Bytes()
}

func (n *Node) ToByteArray() []byte {
	return n.toByteBuffer("").Bytes()
}

func (n *Node) toByteBuffer(prefix string) *bytes.Buffer {
	writer := bytes.NewBufferString(prefix)
	if len(n.Content) == 0 && len(n.Children) == 0 {
		writer.WriteString("<" + n.Name + " " + n.getArgsString() + "/>")
		return writer
	}
	writer.WriteString("<" + n.Name + " " + n.getArgsString() + ">")
	if len(n.Content) > 0 {
		writer.Write(n.Content)
	}
	for _, c := range n.Children {
		writer.Write(c.ToByteArray())
	}
	writer.WriteString("</" + n.Name + ">")
	return writer
}

func (n *Node) getArgsString() string {
	if len(n.Attrs) == 0 {
		return ""
	}
	writer := bytes.NewBufferString("")
	for _, a := range n.Attrs {
		writer.WriteString(a.Name + "=\"" + a.Value + "\" ")
	}
	return writer.String()
}
