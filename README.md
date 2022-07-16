**argsv.go**(Command line argument parser)

I got interested in GO programming language and to learn it, I started working on this project. 


__Example__
```
package main

import (
	"fmt"
	"os"
)

var command = "?,/?,h,-h,/h,help,-help,--help(Displays help screen)#v,-v,/v,version,-version,--version(Displays version number)#d,-d,/d,dir,--dir(List files)"

func main() {

	list := LinkedList{0, nil, nil}
	obj := Parser{command, defaultIndex}
	obj.parser()

	for i := defaultIndex; i <= len(os.Args[1:]); i++ {
		for j := defaultIndex; j <= int(obj.getN()); j++ {
			for k := defaultIndex; k <= int(obj.getNOptions(uint32(j))); k++ {
				if os.Args[i] == obj.getNthCommandNthOption(uint32(j), uint32(k)) {
					list.append(uint32(i), uint32(j), uint32(k))
				}
			}
		}
	}	
	list.update()

	ll := list.find("d", obj)
	if ll.length > 0 {
		fmt.Println("Found...", ll.length)
		node := ll.head
		for i := uint32(0); i < uint32(ll.length); i++ {
			fmt.Println(node.i, " ", node.c, " ", node.o, " ", node.n)
			node = node.next
		}
	} else {
		fmt.Println("Not found...")
	}
}

```


