package core

type Parser struct {
	operators []byte
}

func Parse(source string) interface{} {
	parser := NewParser(source)
	return parser.Parse()
}
