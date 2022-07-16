/*
 * linkedlist.go
 */

// https://dev.to/divshekhar/golang-linked-list-data-structure-h20

package main

import (
	"fmt"
	"os"
)

type Node struct {
	i    uint32 // Index, index into command line arguments
	c    uint32 // Command, the whole command string with all options and help string
	o    uint32 // Option, a single command line option
	n    uint32 // argc for this command line option
	next *Node
	prev *Node
}

type LinkedList struct {
	length uint32

	head *Node
	tail *Node
}

// Receiver object is by reference
func (l *LinkedList) append(i uint32, j uint32, k uint32) bool {

	if l.len() == 0 {
		l.head = &Node{i, j, k, 0, nil, nil}
		l.tail = l.head
	} else {
		l.tail.next = &Node{i, j, k, 0, nil, nil}
		l.tail.next.prev = l.tail
		l.tail = l.tail.next
	}

	l.length++

	return true
}

// Receiver object is pass by value
func (l LinkedList) display(p Parser) {

	node := l.head

	for node != nil {

		fmt.Println("argv index = ", node.i, " Command string = ", node.c, " Command line option = ", node.o, " argc = ", node.n)
		fmt.Println("Command = ", p.getNthCommand(node.c), "Option -> ", p.getNthCommandNthOption(node.c, node.o))

		node = node.next
	}
}

// Receiver object is pass by value
//func (l LinkedList) find(c string, p Parser) *Node {

func (l LinkedList) find(c string, p Parser) *LinkedList {

	node := l.head

	ll := &LinkedList{0, nil, nil}

	retHead := &Node{0, 0, 0, 0, nil, nil}
	ret := retHead

	ll.head = ret

	for node != nil {
		for i := uint32(defaultIndex); i <= p.getNOptions(node.c); i++ {
			if p.getNthCommandNthOption(node.c, i) == c {

				//fmt.Println(ret.i, " ", ret.c, " ", ret.o, " ", ret.n)

				if ret.i == 0 && ret.c == 0 && ret.o == 0 {
					//ret = Node{node.i, node.c, node.o, node.n, nil, nil}
					ret.i = node.i
					ret.c = node.c
					ret.o = node.o
					ret.n = node.n
					ret.next = nil
					ret.prev = nil

					//fmt.Println("Yes...")
				} else {
					ret.next = &Node{node.i, node.c, node.o, node.n, nil, ret}
					ret.next.prev = ret
					ret = ret.next

					//fmt.Println("Hopla...")
				}

				ll.length++
			}
		}
		node = node.next
	}

	/*
		for ret.prev != nil {
			ret = ret.prev
		}
	*/

	ll.tail = ret

	//return retHead
	return ll
}

// Receiver object is pass by value
func (l LinkedList) len() uint32 {

	return l.length
}

// Receiver object is pass by value
func (l LinkedList) update() {

	node := l.head

	for node != nil {

		if node.next != nil {
			node.n = node.next.i - node.i
		} else {
			node.n = uint32(len(os.Args[0:])) - node.i
		}

		node = node.next
	}
}
