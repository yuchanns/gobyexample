package markdown

import (
	"gopkg.in/russross/blackfriday.v2"
)

func MdRun(input []byte) []byte {
	output := blackfriday.Run(input)

	return output
}

func MdParse(input []byte) *blackfriday.Node {
	r := blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{
		Flags: blackfriday.CommonHTMLFlags,
	})
	optList := []blackfriday.Option{
		blackfriday.WithRenderer(r),
		blackfriday.WithExtensions(blackfriday.CommonExtensions),
	}
	parse := blackfriday.New(optList...)

	ast := parse.Parse(input)

	return ast
}
