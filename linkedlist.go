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
func (l LinkedList) find(c string, p Parser) bool {

	node := l.head

	for node != nil {
		for i := uint32(defaultIndex); i <= p.getNOptions(node.c); i++ {
			if p.getNthCommandNthOption(node.c, i) == c {
				return true
			}
		}
		node = node.next
	}

	return false
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
