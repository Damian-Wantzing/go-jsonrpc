package jsonrpc

import (
	"testing"
)

func TestUnmarshalValidIntoRequest(t *testing.T) {
	json := "{\"jsonrpc\":\"2.0\",\"method\":\"rpc.METHOD\",\"params\":[1,2,3],\"id\":1}"

	req, err := ParseRequest([]byte(json))

	if err != nil {
		t.Fatalf("Error parsing request: %s", err)
	}

	if req.Jsonrpc != "2.0" {
		t.Fatal("Invalid JSON-RPC version")
	}

	if req.Method != "rpc.METHOD" {
		t.Fatal("Invalid method")
	}

	if req.ID == nil {
		t.Fatal("Invalid ID")
	}

	if (*req.ID).(float64) != 1 {
		t.Fatal("Invalid ID")
	}

	if req.Params == nil {
		t.Fatal("Invalid params")
	}

	if len(req.Params.(ArrayParams)) != 3 {
		t.Error("Invalid params")
	}

	if req.IsNotification() {
		t.Error("Invalid notification")
	}

	json = "{\"jsonrpc\":\"2.0\",\"method\":\"rpc.METHOD\",\"params\":{\"arg1\": \"val1\", \"arg2\": \"val2\"},\"id\":1}"

	req, err = ParseRequest([]byte(json))
	if err != nil {
		t.Fatalf("Error parsing request: %s", err)
	}

	if req.Params == nil {
		t.Fatal("Invalid params")
	}

	if len(req.Params.(MapParams)) != 2 {
		t.Fatal("Invalid params")
	}
}

func TestIncorrectVersion(t *testing.T) {
	json := "{\"jsonrpc\":\"1.0\",\"method\":\"rpc.METHOD\",\"params\":[1,2,3],\"id\":1}"

	_, err := ParseRequest([]byte(json))

	if err == nil {
		t.Fatal("Error expected due to incorrect version")
	}
}

func TestNoParams(t *testing.T) {
	json := "{\"jsonrpc\":\"2.0\",\"method\":\"rpc.METHOD\",\"id\":1}"

	req, err := ParseRequest([]byte(json))

	if err != nil {
		t.Fatalf("Error parsing request: %s", err)
	}

	if req.Params != nil {
		t.Fatal("Invalid params")
	}
}

func TestNoIDIsNotification(t *testing.T) {
	json := "{\"jsonrpc\":\"2.0\",\"method\":\"rpc.METHOD\"}"

	req, err := ParseRequest([]byte(json))

	if err != nil {
		t.Fatalf("Error parsing request: %s", err)
	}

	if !req.IsNotification() {
		t.Fatal("Not a notification when no ID is present")
	}
}
