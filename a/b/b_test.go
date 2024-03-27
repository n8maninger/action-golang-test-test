package b

import (
	"testing"
	"time"
)

func TestB(t *testing.T) {
	t.Log("TestB")
}

func TestSlow(t *testing.T) {
	t.Log("TestSlow")
	time.Sleep(10 * time.Second)
}

func TestFailure(t *testing.T) {
	t.Log("Foo")
	t.Fatal("bar")
}
