/*
	constants.go
	Q@khaa.pk
 */

package main

/*
  This file defines the tokens used for parsing commands and their arguments.
  You can modify these constants to change the delimiters for:
  - Command argument parsing (currently '#')
  - Command separation (currently ',')
  - Help string start (currently '(')
  - Help string end (currently ')')

  Modify these as needed to suit your application's requirements.
*/

const commandArgumentParserToken rune = '#'
const commandParserToken rune = ','
const commandHelpStringParserTokenEnd rune = ')'
const commandHelpStringParserTokenStart rune = '('

const defaultIndex = 1
