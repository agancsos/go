package main

type test1Action struct {
	Script              string
}

func (x test1Action) InvokeAction() (error, map[string]interface{}) {
	return nil, map[string]interface{}{"script": x.Script,};
}

var Test1Action test1Action;

