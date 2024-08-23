package jsonrpc

import "io"

type JsonRPC struct {
	methods Methods
}

// HandleRequest handles an incoming JSON-RPC request
// and returns the response from the called method
// via the provided writer
func (j *JsonRPC) HandleRequest(writer io.Writer, json []byte) {

}
