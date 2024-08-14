package jsonrpc

// A Request object is used to invoke a method on a remote server
// https://www.jsonrpc.org/specification#request_object
type Request struct {
	Jsonrpc string         `json:"jsonrpc"` // The version of the JSON-RPC protocol, which must always be exactly "2.0"
	Method  string         `json:"method"`  // A string containing the name of the method to be invoked. Must start with "rpc.METHOD"
	Params  map[string]any `json:"params"`  // The parameters to pass to the method
	Id      any            `json:"id"`      // An identifier specified by the client. Must be int, string or null. If omitted, the request is a notification.
}

// A Response object is used to return the result of a method invocation
// https://www.jsonrpc.org/specification#response_object
type Response struct {
	Jsonrpc string `json:"jsonrpc"` // The version of the JSON-RPC protocol, which must always be exactly "2.0"
	Result  any    `json:"result"`  // The result of the call, which must be present on success or absent on error
	Error   Error  `json:"error"`   // An error object if there was an error invoking the method
}

// An Error object is used to return an error in the response
// These are the reserved error codes, which may not be used.
// Other error codes are available for application-specific errors.
//
//	code	            message	            meaning
//	-32700	            Parse error	        Invalid JSON was received by the server.
//	                                        An error occurred on the server while parsing the JSON text.
//	-32600	            Invalid Request	    The JSON sent is not a valid Request object.
//	-32601	            Method not found	The method does not exist / is not available.
//	-32602	            Invalid params	    Invalid method parameter(s).
//	-32603	            Internal error	    Internal JSON-RPC error.
//	-32000 to -32099	Server error	    Reserved for implementation-defined server-errors.
//
// https://www.jsonrpc.org/specification#error_object
type Error struct {
	Code    int    `json:"code"`    // A number that indicates the error type that occurred
	Message string `json:"message"` // A string providing a short description of the error
	Data    any    `json:"data"`    // A Primitive or Structured value that contains additional information about the error and may be omitted
}
