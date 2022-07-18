package main

import (
	"os"
)

func argsvGo(c string, a string) LinkedList {

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

	//ll := list.find(a, obj)

	//return *ll

	return *list.find(a, obj)
}
