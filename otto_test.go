package otto

import (
	"fmt"
	"testing"
)

var called chan bool

func init() {
	called = make(chan bool, 1)
}

func FakeJobHandler(m *Message) {
	fmt.Println("worker started...")
	called <- true
	defer close(called)
	fmt.Println("worker done!")
}

func TestInit(t *testing.T) {
	o := New(&Config{})

	o.Register(&Worker{
		Queue:       "test_queue",
		Handler:     FakeJobHandler,
		Concurrency: 10,
	})

	expected := true
	got := <-called

	if got != expected {
		t.Fatalf("Expected %v, Got %v", expected, got)
	}
}
