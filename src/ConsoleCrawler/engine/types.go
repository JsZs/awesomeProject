package engine

type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

type ParserFunc func(contents []byte, url string) ParseResult

type Item struct {
	Id      string //去重时使用
	Url     string
	Type    string
	Payload interface{}
}
type Request struct {
	Url string
	//ParserFunc func([] byte) ParseResult
	Parser Parser
}

type ParseResult struct {
	Requests []Request
	//Items [] interface{}
	Items []Item
}

type NilParser struct{}

func (NilParser) Parse([]byte, string) ParseResult {
	return ParseResult{}
}

func (NilParser) Serialize() (name string, args interface{}) {
	return "NilParser", nil
}

type FuncParser struct {
	parser ParseResult
	name   string
}

func (f *FuncParser) Parse(contents []byte, url string) ParseResult {
	return f.parser(contents, url)
}
func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}
func NewFuncParser(
	p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}
