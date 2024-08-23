package jsonrpc

import (
	"encoding/json"
	"errors"
)

// Params is an interface that supports either map or array
// parameters for a Request
type Params interface {
	Type() string
}

// ArrayParams is used when the params for the called
// method are provided as a (ordered) array
type ArrayParams []any

// Type returns the type of the params
func (p ArrayParams) Type() string {
	return "array"
}

// MapParams is used when the params for the called
// method are provided as a map (key-value pair)
type MapParams map[string]any

// Type returns the type of the params
func (p MapParams) Type() string {
	return "map"
}

// A Request object is used to invoke a method on a remote server
// https://www.jsonrpc.org/specification#request_object
type Request struct {
	Jsonrpc string // The version of the JSON-RPC protocol, which must always be exactly "2.0"
	Method  string // A string containing the name of the method to be invoked. Must start with "rpc.METHOD"
	Params  Params // The parameters to pass to the method
	ID      *any   // An identifier specified by the client. Must be int, string or null. If omitted, the request is a notification.
}

// Custom unmarhshal method to parse a raw JSON into the correct
// Request object
func (r *Request) UnmarshalJSON(data []byte) error {

	// A raw request that we can use to determine the params
	type rawRequest struct {
		Jsonrpc string `json:"jsonrpc"`
		Method  string `json:"method"`
		ID      *any   `json:"id,omitempty"`
		Params  *any   `json:"params,omitempty"`
	}

	var raw rawRequest
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	if raw.Jsonrpc != "2.0" {
		return errors.New("Invalid JSON-RPC version")
	}

	r.Jsonrpc = raw.Jsonrpc
	r.Method = raw.Method
	r.ID = raw.ID

	if raw.Params == nil {
		return nil
	}

	if params, ok := (*raw.Params).(map[string]any); ok {
		params := MapParams(params)
		r.Params = params
		return nil
	}

	if params, ok := (*raw.Params).([]interface{}); ok {
		params := ArrayParams(params)
		r.Params = params
		return nil
	}

	return errors.New("Invalid params type")
}

// Returns whether the request is a notification
func (r *Request) IsNotification() bool {
	return r.ID == nil
}

// ParseRequest parses a raw json request into
// a request object
func ParseRequest(data []byte) (Request, error) {
	var req Request
	err := json.Unmarshal(data, &req)
	return req, err
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
