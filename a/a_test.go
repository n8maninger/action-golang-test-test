package a

import (
	"testing"
	"time"
)

func TestA(t *testing.T) {
	t.Log("TestA")
}

func TestSlow(t *testing.T) {
	t.Log("TestSlow")
	time.Sleep(10 * time.Second)
}

func TestFailure(t *testing.T) {
	t.Log("Foo")
	t.Fatal("bar")
}
