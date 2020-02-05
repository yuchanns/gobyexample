package markdown

import "gopkg.in/russross/blackfriday.v2"

func MdRun(input []byte) []byte {
	output := blackfriday.Run(input)

	return output
}
