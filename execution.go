package jsonrpc

// A single parameter for an Executable method
type parameter struct {
	Name string
	Type string
}

// This is a wrapper around an executable method
// that is exposed to the JSONRPC protocol
type Executable struct {
	name   string
	exec   func(params map[string]interface{}) (interface{}, error)
	params []parameter
}

// Execute an executable method
func (e *Executable) execute(request Request) (interface{}, error) {
	// TODO: we probably cannot pass the params as is
	// since they can be passed in two ways: by name or by position
	return e.exec(request.Params)
}

// The Executor hold the methods that will be executed
// when specifically requested by the JSONRPC protocol
// if you want to expose a new method to JSONRPC, this
// is where you add it
type Executor struct {
	methods map[string]Executable
}

// Creates and returns a new Executor
func New() Executor {
	return Executor{
		methods: make(map[string]Executable),
	}
}

// Add a method to the Executor
// This will expose the method to the JSONRPC protocol
func (e *Executor) Add(method Executable) {
	e.methods[method.name] = method
}

// Remove a method from the executor
// This will hide the method from the JSONRPC protocol
func (e *Executor) Remove(name string) {
	delete(e.methods, name)
}

// Execute a method from a JSONRPC message
func (e *Executor) execute(request Request) (*Response, *Error) {
	return nil, nil
}
