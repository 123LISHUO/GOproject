package main

import (
	"fmt"
)

type Node struct {
	data int
	prev *Node
	next *Node
}

func InsertNode(root *Node, data int) *Node {
	newNode := &Node{data: data}
	if root == nil {
		return newNode
	}
	curr := root
	for curr.next != nil && curr.data < data {
		curr = curr.next
	}
	if curr.data >= data {
		newNode.next = curr
		newNode.prev = curr.prev
		if curr.prev != nil {
			curr.prev.next = newNode
		} else {
			root = newNode
		}
		curr.prev = newNode
	} else {
		newNode.prev = curr
		curr.next = newNode
	}
	return root
}

func RangleLink(root *Node) {
	tmpnode := root
	for {

		if tmpnode == nil {
			break
		}
		fmt.Println(tmpnode.data)
		tmpnode = tmpnode.next

	}
}

func DeleteNode(root *Node, data int) *Node {
    if root == nil {
        return nil
    }
    curr := root
    for curr != nil && curr.data != data {
        curr = curr.next
    }
    if curr == nil {
        return root
    }
    if curr.prev != nil {
        curr.prev.next = curr.next
    } else {
        root = curr.next
    }
    if curr.next != nil {
        curr.next.prev = curr.prev
    }
    return root
}


func main() {
	n1 := &Node{
		data: 1,
		prev: nil,
	}

	n2 := &Node{
		data: 2,
		prev: n1,
	}

	n3 := &Node{
		data: 3,
		prev: n2,
	}
	n4 := &Node{
		data: 4,
		prev: n3,
	}
	n6 := &Node{
		data: 6,
		prev: n4,
	}

	n1.next = n2
	n2.next = n3
	n3.next = n4
	n4.next = n6
	n6.next = nil

	root := InsertNode(n1, 1)
	root = DeleteNode(root, 1)
	RangleLink(root)
}
