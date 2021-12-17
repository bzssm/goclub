package decorator

type Context struct {
	// InputParams
	InputParams map[string]interface{}
	// OutputParams
	OutputParams map[string]interface{}
	// FuncName is outer caller's name
	FuncName string
	// Keys are used for some storage
	Keys map[string]interface{}
}

func NewContext() *Context {
	return &Context{
		InputParams:  make(map[string]interface{}),
		OutputParams: make(map[string]interface{}),
		Keys:         make(map[string]interface{}),
	}
}

func (c *Context) Reset() {
	c.InputParams = make(map[string]interface{})
	c.OutputParams = make(map[string]interface{})
	c.FuncName = ""
	c.Keys = make(map[string]interface{})
}

type Handler func(c *Context)
