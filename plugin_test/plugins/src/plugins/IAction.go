package plugins

type IAction interface {
	InvokeAction()     (error, map[string]interface{})
}

