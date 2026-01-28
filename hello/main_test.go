package hello

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, world!"
	
	got := Hello()

	if got != want {
		t.Errorf("Hello() = %v, want %v", got, want)
	}
}
