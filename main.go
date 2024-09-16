/*
	 main.go
	 Q@khaa.pk
 */

package main

import (
		 "fmt"
		 "os"
)

var command = "?,/?,h,-h,/h,help,-help,--help(Displays help screen)#v,-v,/v,version,-version,--version(Displays version number)#d,-d,/d,dir,--dir(List files)"

func main() {

	ll := argsvGo(command, "dir")

	if ll.length > 0 {
		fmt.Println("Command found, number of instances = ", ll.length)
		node := ll.head
		for i := uint32(0); i < ll.length; i++ {
			fmt.Println( "Index into command string = ", node.c, 
						 "\nIndex into argv = ", node.i, "\nargc of this command = ", node.n, "\nIndex into the command string for one praticulr name used = ", node.o, "\n and it is \"", os.Args[node.i] , "\"")
			
			fmt.Println("The command and its arguments are...") 			 
			for j := uint32(0); j < node.n; j++ {
				fmt.Println(os.Args[node.i + j])
			}

			fmt.Println("Help String = " + argsvGoHelp(node.c) + "\n")
			node = node.next			
		}
	}

	fmt.Println("\nNumber of common arguments = ", ll.commonArgs)
	for i := uint32(0); i < ll.commonArgs; i++ {
		fmt.Println(os.Args[i])
	}
}
