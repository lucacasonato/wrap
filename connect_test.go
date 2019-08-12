package wrap_test

import (
	"testing"
	"time"

	"github.com/lucacasonato/wrap"
)

func connect() (*wrap.Client, error) {
	return wrap.Connect("mongodb://localhost:27017", 2*time.Second)
}

func TestConnect(t *testing.T) {
	client, err := connect()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(client)
}

func TestConnectFail(t *testing.T) {
	client, err := wrap.Connect("mongodb://brokenhost:0", 2*time.Second)
	if err == nil {
		t.Fatal("connect did not error on an invalid connection")
	}

	t.Log(client)
}