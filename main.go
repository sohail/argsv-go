/*
 * main.go
 */

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

	/*
		fmt.Println(obj.getN())


		fmt.Println(obj.getNthCommand(0))
		fmt.Println(obj.getNthCommand(3))
		fmt.Println(obj.getNthCommand(1))
		fmt.Println(obj.getNthCommand(2))

		fmt.Println(obj.getNthCommandHelpString(1))
		fmt.Println(obj.getNthCommandHelpString(2))
		fmt.Println(obj.getNthCommandHelpString(3))

		fmt.Println(obj.getNthCommandOptions(1))
		fmt.Println(obj.getNthCommandOptions(2))
		fmt.Println(obj.getNthCommandOptions(3))
	*/

	//fmt.Println(obj.getNOptions(3))
	//fmt.Println(obj.getNthCommandNthOption(1, 7))

	/*
		for i := defaultIndex; i <= int(obj.getN()); i++ {
			fmt.Println("--> ", obj.getNthCommandOptions(uint32(i)))
			for j := defaultIndex; j <= int(obj.getNOptions(uint32(i))); j++ {
				fmt.Println("---> ", obj.getNthCommandNthOption(uint32(i), uint32(j)))
			}
		}

		for i := defaultIndex; i <= len(os.Args[1:]); i++ {
			fmt.Println(os.Args[i])
		}
	*/

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

	//fmt.Println("--> ", list.len())

	//fmt.Println(len(os.Args[1:]))

	//fmt.Println(obj.getNthCommand(4))

	list.display(obj)

	//fmt.Println("-> ", os.Args[1:list.maxCommonArgs()])

	if list.find("d", obj) {
		fmt.Println("Command found")
	} else {
		fmt.Sprintln("Command not found")
	}

}
