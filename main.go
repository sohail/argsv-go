/*
 * main.go
 */

package main

import "fmt"

var command = "?,/?,h,-h,/h,help,-help,--help(Displays help screen)#v,-v,/v,version,-version,--version(Displays version number)#d,-d,/d,dir,--dir(List files)"

func main() {

	ll := argsvGo(command, "dir")

	if ll.length > 0 {
		fmt.Println("Command found, number of instances = ", ll.length)
		node := ll.head
		for i := uint32(0); i < ll.length; i++ {
			fmt.Println(node.c, " ", node.i, " ", node.n, " ", node.o)
			node = node.next
		}
	}

	fmt.Println("--> ", ll.commonArgs)
}
