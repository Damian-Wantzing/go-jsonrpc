package jsonrpc

// A single parameter for a Function
type parameter struct {
	Name string
	Type string
}

// This is a wrapper around a Function
// that is exposed to the JSONRPC protocol
type Function struct {
	name   string
	exec   func(params map[string]interface{}) (interface{}, error)
	params []parameter
}

// Execute a Function
func (m *Function) execute(request Request) (interface{}, error) {
	// TODO: we probably cannot pass the params as is
	// since they can be passed in two ways: by name or by position
	// return e.exec(request.Params)

	return nil, nil
}

// The Methods struct hold the methods that will be executed
// when specifically requested by the JSONRPC protocol
// if you want to expose a new method to JSONRPC, this
// is where you add it
type Methods struct {
	functions map[string]Function
}

// Creates and returns a new Methods
func New() Methods {
	return Methods{
		functions: make(map[string]Function),
	}
}

// Add a method to the Methods
// This will expose the method to the JSONRPC protocol
func (e *Methods) Add(function Function) {
	e.functions[function.name] = function
}

// Remove a method from the executor
// This will hide the method from the JSONRPC protocol
func (e *Methods) Remove(name string) {
	delete(e.functions, name)
}

// Execute a method from a JSONRPC message
func (e *Methods) execute(request Request) (*Response, *Error) {
	return nil, nil
}
