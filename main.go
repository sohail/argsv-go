/*
	 main.go
	 Q@khaa.pk
 */

package main

import (
		 "fmt"
		 "os"
)

/*
	Each command group is separated by #,
	and within each group, 
	different aliases for the same command are separated by commas.
	The text in parentheses provides help information for each command.

	?, /h, -h, --help: These are aliases for the "help" command.
	v, -v, version: Aliases for version command.
	d, -d, dir: These are for directory listing.
	
	Note:
	The delimiters used for parsing commands and arguments in this file can be customized.
	To modify the delimiters (e.g., for command argument parsing, command separation, or help strings),
	please refer to the `constants.go` file where these tokens are defined.
 */
var command = "?,/?,h,-h,/h,help,-help,--help(Displays help screen)#v,-v,/v,version,-version,--version(Displays version number)#d,-d,/d,dir,--dir(List files)"

func main() {

	// Call to argsvGo returns a linked list (`ll`) containing all instances of the "dir" command found in the input.
	ll := argsvGo(command, "dir")

	// The `ll.length` property indicates the number of times the "dir" command was detected.
	if ll.length > 0 {
		fmt.Println("Command found, number of instances = ", ll.length)
		// Each node contains details about a specific occurrence of the command.
		/*
			`node.c`: Index into the command string.
			`node.i`: Index into the argument list (`argv`).
			`node.n`: The number of arguments associated with this instance of the command.
			`node.o`: The index into the command string for the specific alias used.
		 */
		node := ll.head
		// The loop iterates through the linked list, printing these details for each occurrence of the "dir" command.
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
	else {
		fmt.Println("Not a single instance of \"dir\" found in input.")
	}

	/*
		This code block handles the printing of common arguments (arguments that are not associated with any specific command).
		
		- `ll.commonArgs` holds the number of arguments that are shared or common across all commands.
		- The loop iterates through these common arguments (from index 0 up to `ll.commonArgs`), printing each one.
		- `os.Args[i]` accesses the command-line arguments, where `i` is the index of the argument.
	 */
	fmt.Println("\nNumber of common arguments = ", ll.commonArgs)
	for i := uint32(0); i < ll.commonArgs; i++ {
		fmt.Println(os.Args[i])
	}
}
