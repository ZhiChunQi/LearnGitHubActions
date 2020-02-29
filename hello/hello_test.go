package hello

import "testing"

func TestHello(t *testing.T) {
	result := Hello()
	if result != "Hello Actions" {
		t.Error("fail")
		return
	}
	t.Log("success")
}
