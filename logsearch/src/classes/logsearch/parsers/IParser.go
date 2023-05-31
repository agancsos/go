package parsers

type IParser interface {
	Parse(raw string)  (error, interface{})
	Unparse(obj interface{}) (error, string)
}

type DummyParser struct {}
func (x DummyParser) Parse(raw string)  (error, interface{}) {
	return nil, "";
}
func (x DummyParser) Unparse(obj interface{}) (error, string) {
	return nil, "";
}

