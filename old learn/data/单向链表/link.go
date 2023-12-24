package main
// package main

import (
	"fmt"
)
// 定义链表结构体
type Liknode struct {
	data int
	next *Liknode
}
func main() {
	//在链表里面写入数据
	n1 := &Liknode{}
	n2 := &Liknode{}
	n3 := &Liknode{}
	n4 := &Liknode{}
	n6 := &Liknode{}

	n1.data = 1
	n2.data = 2
	n3.data = 3
	n4.data = 4
	n6.data = 6

	n1.next = n2
	n2.next = n3
	n3.next = n4
	n4.next = n6
	n6.next = nil

	// n5 := &Liknode{
	// 	data: 5,
	// 	next: nil,
	// }

	Printlink(n1)
	n1 = DeleteNode(n1, 1)
	n1 = DeleteNode(n1, 2)
	n1 = DeleteNode(n1, 3)
	Printlink(n1)

	
}

// 在单向链表中中插入node
func InsertNode(n1 *Liknode, n5 *Liknode) {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("error:", err)
		}
	}()

	//遍历单向链表，插入node
	tmpnode := n1
	for {
		if tmpnode != nil {
			// fmt.Println(tmpnode.data)
			if n5.data > tmpnode.data {
				if tmpnode.next == nil {
					tmpnode.next = n5
					n5.next = nil
				} else {
					if n5.data <= tmpnode.next.data {
						n5.next = tmpnode.next
						tmpnode.next = n5
					}
				}
			}
		} else {
			break
		}
		tmpnode = tmpnode.next
	}
	fmt.Println("插入成功！",n5.data)
}

// 删除单向链表中的node
func DeleteNode(root *Liknode, data int) *Liknode{

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("error:", err)
		}
	}()

	tmpnode := root
	// 删除链表的第一个node
	if root != nil && root.data == data {
		if root.next == nil {
			return nil
		}
		right := tmpnode.next
		tmpnode.next = nil
		return right
	}
	// 删除链表其他的node
	for {
		if tmpnode != nil {
			break
		}

		right := tmpnode.next

		if right.data == data {
			tmpnode.next = right.next
			right.next = nil
			return root
		}

		tmpnode = tmpnode.next
	}
	fmt.Println("删除成功！",data)
	return root
}

// 打印链表中的数据
func Printlink(n1 *Liknode) {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("error:", err)
		}
	}()

	tmpnode := n1
	for {
		fmt.Println(tmpnode.data)
		if tmpnode.next == nil {
			break
		}
		
		tmpnode = tmpnode.next
	}

	// for tmpnode:= n1;tmpnode != nil;tmpnode = tmpnode.next{
	// 	fmt.Println(tmpnode.data)
	// }
	// fmt.Println("打印成功！")
	// return 
}
