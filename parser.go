/*
	parser.go
	Q@khaa.pk
 */

package main

import (
	"os"
	"strings"
)

type Parser struct {
	command string
	n       uint32 // Number of commands
}

// Receiver object is pass by value
func (p Parser) getN() uint32 {
	return p.n
}

// Receiver object is pass by value
func (p Parser) getNOptions(i uint32) uint32 {

	if i <= p.n {
		var j uint32 = defaultIndex
		var o string = p.getNthCommandOptions(i)

		if o != "" {
			_ = strings.FieldsFunc(o, func(r rune) bool {

				if r == commandParserToken {
					j++
					return true
				}

				return false
			})

			return j
		}
	}

	return 0
}

// Receiver object is pass by value
func (p Parser) getNthCommand(i uint32) string {

	if i == 0 || i > p.n {
		return ""
	}

	cuttingByTwo := strings.FieldsFunc(p.command, func(r rune) bool {

		if r == commandArgumentParserToken {
			return true
		}

		return false
	})

	return cuttingByTwo[i-1]
}

// Receiver object is pass by value
func (p Parser) getNthCommandNthOption(i uint32, j uint32) string {

	var o string = p.getNthCommandOptions(i)
	var k uint32 = defaultIndex

	if o != "" {

		cuttingByTwo := strings.FieldsFunc(o, func(r rune) bool {

			if r == commandParserToken {
				k++
				return true
			}

			return false
		})

		if k-1 >= j-1 {
			return cuttingByTwo[j-1]
		}
	}

	return ""
}

// Receiver object is pass by value
func (p Parser) getNthCommandOptions(i uint32) string {

	var c string = p.getNthCommand(i)

	if c != "" {

		cuttingByTwo := strings.FieldsFunc(c, func(r rune) bool {

			if r == commandHelpStringParserTokenStart {
				return true
			}

			return false
		})

		return cuttingByTwo[0]
	}

	return ""
}

// Receiver object is pass by value
func (p Parser) getNthCommandHelpString(i uint32) string {

	var c string = p.getNthCommand(i)

	//fmt.Println(c)
	
	if c != "" {

		cuttingByTwo := strings.FieldsFunc(c, func(r rune) bool {

			if r == commandHelpStringParserTokenStart {
				return true
			}

			return false
		})

		cuttingByTwo = strings.FieldsFunc(cuttingByTwo[1], func(r rune) bool {

			if r == commandHelpStringParserTokenEnd {
				return true
			}

			return false
		})

		return cuttingByTwo[0]
	}

	return ""
}

// Receiver object is pass by reference
func (p *Parser) parser() string {

	cuttingByTwo := strings.FieldsFunc(p.command, func(r rune) bool {

		if r == commandArgumentParserToken {
			p.n++
			return true
		}

		return false
	})

	return cuttingByTwo[1]
}

func (l LinkedList) maxCommonArgs() uint32 {

	var n uint32 = uint32(len(os.Args[0:]))

	node := l.head

	for node != nil {

		if n > node.i {
			n = node.i
		}

		node = node.next
	}

	return n
}
